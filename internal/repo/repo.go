package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// ErrNotFound is returned when no entity with a such id exists in the repo
var ErrNotFound = errors.New("entity already exists")

// ErrNotFound is returned when given limit and offset are invalid
var ErrOutOfRange = errors.New("limit and offset are out of range")

type Repo interface {
	AddEntities(ctx context.Context, entities []models.Entity) error
	RemoveEntity(ctx context.Context, entityId uint64) error
	DescribeEntity(ctx context.Context, entityId uint64) (*models.Entity, error)
	ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error)
}

// type InMemoryFeedbackRepo struct {
// 	mutex sync.RWMutex
// 	data  []models.Feedback
// }

// type InMemoryProposalRepo struct {
// 	mutex sync.RWMutex
// 	data  models.Proposal
// }
