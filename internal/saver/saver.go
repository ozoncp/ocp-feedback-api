package saver

import (
	"errors"
	"log"

	"github.com/ozoncp/ocp-feedback-api/internal/alarmer"
	"github.com/ozoncp/ocp-feedback-api/internal/flusher"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

type Policy int

const (
	DropAll Policy = iota // drop all data
	DropOne               // drop only the oldest data
)

type Saver interface {
	Save(entity models.Entity)
}

type void struct{}

type saver struct {
	policy     Policy
	entitiesCh chan models.Entity
	entities   []models.Entity
	alarmer    alarmer.Alarmer
	flusher    flusher.Flusher
	done       chan void
}

func New(capacity int, policy Policy,
	alarmer alarmer.Alarmer, flusher flusher.Flusher) (*saver, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than 0")
	}
	if alarmer == nil {
		return nil, errors.New("alarmer cannot be nil")
	}
	if flusher == nil {
		return nil, errors.New("flusher cannot be nil")
	}

	return &saver{
		policy:     policy,
		entitiesCh: make(chan models.Entity),
		entities:   make([]models.Entity, 0, capacity),
		alarmer:    alarmer,
		flusher:    flusher,
		done:       make(chan void),
	}, nil
}

func (s *saver) Close() {
	close(s.done)
}

func (s *saver) Save(entity models.Entity) {
	s.entitiesCh <- entity
}

func (s *saver) Init() {
	go func() {
		for {
			select {
			case entity := <-s.entitiesCh:
				if len(s.entities) == cap(s.entities) {
					switch s.policy {
					case DropAll:
						s.entities = s.entities[:0]
					case DropOne:
						copy(s.entities[0:], s.entities[1:])
						s.entities = s.entities[:len(s.entities)-1]
					}
				}
				s.entities = append(s.entities, entity)
			case _, ok := <-s.alarmer.Alarm():
				if ok {
					rem, err := s.flusher.Flush(s.entities)
					s.entities = s.entities[0:copy(s.entities[0:], rem)]
					if err != nil {
						log.Printf("failed to save: %v", err)
					}
				}
			case <-s.done:
				if _, err := s.flusher.Flush(s.entities); err != nil {
					log.Printf("failed to save: %v", err)
				}
				return
			}
		}
	}()
}

// 1 2 3 4 5 6
