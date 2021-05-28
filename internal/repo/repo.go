package repo

import "github.com/ozoncp/ocp-feedback-api/internal/models"

type Repo interface {
	AddEntities(entity []models.Entity) error
	RemoveEntity(entityId uint64) error
	DescribeEntity(entityId uint64) (*models.Entity, error)
	ListEntities(limit, offset uint64) ([]models.Entity, error)
}
