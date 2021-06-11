package repo

import (
	"context"
	"errors"
	"sync"

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

// InMemoryFeedbackStore stores feedbacks in memory
type InMemoryFeedbackRepo struct {
	mutex     sync.RWMutex
	data      []*models.Feedback //! slice is used for simplicity
	sequencer uint64
}

// InMemoryProposalRepo stores proposals in memory
type InMemoryProposalRepo struct {
	mutex     sync.RWMutex
	data      []*models.Proposal //! slice is used for simplicity
	sequencer uint64
}

// AddEntities adds feedbacks to the repo and returns ids of inserted objects
func (r *InMemoryFeedbackRepo) AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var ids []uint64

	for i := 0; i < len(entities); i++ {
		f, ok := entities[i].(*models.Feedback)
		if !ok {
			return ids, errors.New("underlying type must be *models.Feedback")
		}
		r.sequencer++
		f.Id = r.sequencer
		ids = append(ids, f.Id)
		r.data = append(r.data, f)
	}
	return ids, nil
}

// RemoveEntity removes feedback from the repo
func (r *InMemoryFeedbackRepo) RemoveEntity(ctx context.Context, entityId uint64) error {
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

// DescribeEntity searches for a feedback with an Id
func (r *InMemoryFeedbackRepo) DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for i := 0; i < len(r.data); i++ {
		if r.data[i].Id == entityId {
			return r.data[i], nil
		}
	}
	return nil, ErrNotFound
}

// ListEntities returns a list of at most 'limit' feedbacks starting from 'offset'
func (r *InMemoryFeedbackRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.data) == 0 {
		return nil, nil
	}

	var feedbacks []models.Entity

	for i := offset; i < offset+limit; i++ {
		if i > uint64(len(r.data)-1) { // check out-of-range
			break
		}
		feedbacks = append(feedbacks, r.data[i])
	}
	return feedbacks, nil
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
