package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// ErrNotFound is returned when no record with a such id exists in the repo
var ErrNotFound = errors.New("record not found")

// flusher and saver don't need the whole Repo interface
type BatchAdder interface {
	AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error)
}

type Repo interface {
	BatchAdder
	RemoveEntity(ctx context.Context, entityId uint64) error
	DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error)
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error)
	UpdateEntity(ctx context.Context, entity models.Entity) error
}
