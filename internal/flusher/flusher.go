package flusher

import (
	"context"
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
)

type Flusher interface {
	Flush(ctx context.Context, entities []models.Entity) ([]models.Entity, error)
}

type flusher struct {
	chunkSize int
	repo      repo.BatchAdder
}

// New returns a new flusher object
func New(chunkSize int, repo repo.BatchAdder) (*flusher, error) {
	if chunkSize < 0 {
		return nil, errors.New("chunk size cannot be negative")
	}
	if repo == nil {
		return nil, errors.New("repo cannot be nil")
	}
	return &flusher{chunkSize, repo}, nil
}

// Flush tries to push given entities into the repo
// Entities that can't be pushed will be returned along with an error
func (f *flusher) Flush(ctx context.Context, entities []models.Entity) ([]models.Entity, error) {
	chunks, err := utils.SplitSlice(entities, f.chunkSize)
	if err != nil {
		return entities, fmt.Errorf("unable to flush: %v", err)
	}

	for i := 0; i < len(chunks); i++ {
		if err := f.repo.AddEntities(ctx, chunks[i]); err != nil {
			return entities[i*f.chunkSize:], fmt.Errorf("unable to flush: %v", err)
		}
	}
	return nil, nil // all entities have been flushed successfully
}
