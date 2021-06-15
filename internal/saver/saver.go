package saver

import (
	"context"
	"errors"
	"log"

	"github.com/ozoncp/ocp-feedback-api/internal/alarmer"
	"github.com/ozoncp/ocp-feedback-api/internal/flusher"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// Policy is a set of actions which should be taken if buffer overflow occurs
type Policy int

const (
	DropAll Policy = iota // drop all the data
	DropOne               // drop only the oldest data
)

// Saver is the interface that allow data to be saved
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

// New returns a new Saver object
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

// WaitClosed waits intil saver is closed
func (s *saver) WaitClosed() {
	<-s.done
}

// Save schedules an entity to be flushed into the repo
func (s *saver) Save(entity models.Entity) {
	s.entitiesCh <- entity
}

// Init starts repeatedly processing incoming events
// Received entities will be handled according to the policy
// Each time Saver is notified by Alarmer, it will try to flush stored entities
// If flushing fails, remaining entities will wait until next Alarmer signal occurs
// or until Close is called
// If Close is called, remaining entities will be flushed to the repo
func (s *saver) Init(ctx context.Context) {
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
					rem, err := s.flusher.Flush(ctx, s.entities)
					s.entities = s.entities[:copy(s.entities, rem)]
					if err != nil {
						log.Printf("failed to save: %v", err)
					}
				}
			case <-ctx.Done():
				if _, err := s.flusher.Flush(ctx, s.entities); err != nil {
					log.Printf("failed to save: %v", err)
				}
				s.done <- void{}
				return
			}
		}
	}()
}
