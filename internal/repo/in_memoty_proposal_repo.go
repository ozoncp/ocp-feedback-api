package repo

import (
	"context"
	"errors"
	"sync"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// NewInMemoryProposalRepo returns a new InMemoryProposalRepo
func NewInMemoryProposalRepo() *InMemoryProposalRepo {
	return &InMemoryProposalRepo{data: make([]*models.Proposal, 0)}
}

// InMemoryProposalRepo stores proposals in memory
type InMemoryProposalRepo struct {
	mutex     sync.RWMutex
	data      []*models.Proposal //! slice is used for simplicity
	sequencer uint64
}

// AddEntities adds proposals to the repo and returns ids of inserted objects
func (r *InMemoryProposalRepo) AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var ids []uint64

	for i := 0; i < len(entities); i++ {
		p, ok := entities[i].(*models.Proposal)
		if !ok {
			return ids, errors.New("underlying type must be *models.Proposal")
		}
		r.sequencer++
		p.Id = r.sequencer
		ids = append(ids, p.Id)
		r.data = append(r.data, p)
	}
	return ids, nil
}

// RemoveEntity removes proposal from the repo
func (r *InMemoryProposalRepo) RemoveEntity(ctx context.Context, entityId uint64) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// !linear algorithm is used for simplicity
	for i := 0; i < len(r.data); i++ {
		if r.data[i].Id == entityId {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// DescribeEntity searches for a proposal with an Id
func (r *InMemoryProposalRepo) DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for i := 0; i < len(r.data); i++ {
		if r.data[i].Id == entityId {
			return r.data[i], nil
		}
	}
	return nil, ErrNotFound
}

// ListEntities returns a list of at most 'limit' proposals starting from 'offset'
func (r *InMemoryProposalRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.data) == 0 {
		return nil, nil
	}

	var proposals []models.Entity

	for i := offset; i < offset+limit; i++ {
		if i > uint64(len(r.data)-1) { // check out-of-range
			break
		}
		proposals = append(proposals, r.data[i])
	}
	return proposals, nil
}
