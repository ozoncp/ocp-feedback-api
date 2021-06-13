package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

type proposalRepo struct {
	db *sqlx.DB
}

func NewProposalRepo(db *sqlx.DB) *proposalRepo {
	return &proposalRepo{db: db}
}

// AddEntities inserts proposals into the database
func (r *proposalRepo) AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %v", err)
	}

	stmt, err := tx.PrepareContext(
		ctx,
		`INSERT INTO reaction.proposal (id, user_id, lesson_id, document_id) VALUES
		 (nextval('reaction.proposal_id_seq'::regclass), $1, $2, $3) RETURNING id;`,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to prepare statement: %v", err)
	}

	defer stmt.Close()

	var ids []uint64 // inserted identifiers

	for i := 0; i < len(entities); i++ {
		p, ok := entities[i].(*models.Proposal)
		if !ok {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("unable to rollback transaction: %v", err)
			}
			return nil, errors.New("underlying type must be *models.Proposal")
		}

		var assignedId uint64
		err = stmt.QueryRowContext(ctx, p.UserId, p.LessonId, p.DocumentId).Scan(&assignedId)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("unable to rollback transaction: %v", err)
			}
			return nil, fmt.Errorf("unable to insert a record: %v", err)
		}
		ids = append(ids, assignedId)
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("unable to rollback transaction: %v", err)
		}
		return nil, fmt.Errorf("unable to commit transaction: %v", err)
	}
	return ids, nil
}

// RemoveEntity removes a proposal from the database
func (r *proposalRepo) RemoveEntity(ctx context.Context, entityId uint64) error {

	// check if record exists
	var dummy uint64
	err := r.db.QueryRowContext(ctx,
		"SELECT 1 FROM reaction.proposal WHERE id=$1;",
		entityId,
	).Scan(&dummy)

	if err == sql.ErrNoRows {
		return errors.New("no such proposal")
	} else if err != nil {
		return fmt.Errorf("unable to remove proposal: %v", err)
	}

	_, err = r.db.ExecContext(ctx,
		"DELETE FROM reaction.proposal WHERE id=$1;",
		entityId,
	)
	if err != nil {
		return fmt.Errorf("unable to remove proposal: %v", err)
	}
	return nil
}

// RemoveEntity retrieves a proposal from the database
func (r *proposalRepo) DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error) {

	pr := &models.Proposal{}
	err := r.db.QueryRowContext(ctx,
		"SELECT id, user_id, lesson_id, document_id FROM reaction.proposal WHERE id=$1;",
		entityId,
	).Scan(&pr.Id, &pr.UserId, &pr.LessonId, &pr.DocumentId)

	if err == sql.ErrNoRows {
		return nil, errors.New("no such proposal")
	} else if err != nil {
		return nil, fmt.Errorf("unable to get proposal: %v", err)
	}
	return pr, nil
}

// RemoveEntity returns a list of at most 'limit' proposals starting from 'offset'
func (r *proposalRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, user_id, lesson_id, document_id FROM reaction.proposal LIMIT $1 OFFSET $2;",
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get proposals: %v", err)
	}
	defer rows.Close()

	var proposals []models.Entity
	for rows.Next() {
		pr := &models.Proposal{}
		err := rows.Scan(&pr.Id, &pr.UserId, &pr.LessonId, &pr.DocumentId)
		if err != nil {
			return nil, fmt.Errorf("unable to get proposals: %v", err)
		}
		proposals = append(proposals, pr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to get proposals: %v", err)
	}
	return proposals, nil
}

// UpdateEntity updates a proposal
func (r *proposalRepo) UpdateEntity(ctx context.Context, entity models.Entity) error {

	p, ok := entity.(*models.Proposal)
	if !ok {
		return errors.New("underlying type must be *models.Proposal")
	}

	// check if record exists
	var dummy uint64
	err := r.db.QueryRowContext(ctx,
		"SELECT 1 FROM reaction.proposal WHERE id=$1;",
		p.Id,
	).Scan(&dummy)

	if err == sql.ErrNoRows {
		return errors.New("no such proposal")
	} else if err != nil {
		return fmt.Errorf("unable to remove feedback: %v", err)
	}

	_, err = r.db.ExecContext(ctx,
		"UPDATE reaction.proposal SET user_id=$1, lesson_id=$2, document_id=$3 WHERE id=$4;",
		p.UserId, p.LessonId, p.DocumentId, p.Id,
	)

	if err == sql.ErrNoRows {
		return errors.New("no such proposal")
	} else if err != nil {
		return fmt.Errorf("unable to update proposal: %v", err)
	}
	return nil
}
