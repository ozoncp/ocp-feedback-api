package saver

import (
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/alarmer"
	"github.com/ozoncp/ocp-feedback-api/internal/flusher"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

type OverflowRule int

const (
	DropAll    OverflowRule = iota // drop all data
	DropOldest                     // drop only the oldest data
)

type Saver interface {
	Save(entity models.Entity)
}

type void struct{}

type saver struct {
	rule    OverflowRule
	stored  chan models.Entity
	alarmer alarmer.Alarmer
	flusher flusher.Flusher
	done    chan void
	errs    chan<- error
}

func New(capacity int,
	rule OverflowRule,
	alarmer alarmer.Alarmer,
	flusher flusher.Flusher,
	errs chan<- error) (*saver, error) {

	if capacity < 0 {
		return nil, errors.New("capacity cannot be negative")
	}
	if alarmer == nil {
		return nil, errors.New("alarmer cannot be nil")
	}
	if flusher == nil {
		return nil, errors.New("flusher cannot be nil")
	}
	if errs == nil {
		return nil, errors.New("errors channel cannot be nil")
	}

	return &saver{
		rule:    rule,
		stored:  make(chan models.Entity, capacity),
		alarmer: alarmer,
		flusher: flusher,
		done:    make(chan void),
		errs:    make(chan<- error),
	}, nil
}

func (s *saver) Close() {
	close(s.done)
}

func (s *saver) Save() {

}
