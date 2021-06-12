package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// ErrNotFound is returned when no entity with a such id exists in the repo
var ErrNotFound = errors.New("entity already exists")

// flusher and saver don't need the whole Repo interface
type BatchAdder interface {
	AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error)
}

type Repo interface {
	BatchAdder
	RemoveEntity(ctx context.Context, entityId uint64) error
	DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error)
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error)
}
