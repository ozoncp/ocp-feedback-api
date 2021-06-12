package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

type feedbackRepo struct {
	db *sqlx.DB
}

func NewFeedbackRepo(db *sqlx.DB) *feedbackRepo {
	return &feedbackRepo{db: db}
}

// AddEntities inserts feedback records into the database
func (r *feedbackRepo) AddEntities(ctx context.Context, entities ...models.Entity) ([]uint64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to start transaction: %v", err)
	}

	stmt, err := tx.PrepareContext(
		ctx,
		`INSERT INTO reaction.feedback (id, user_id, classroom_id, comment) VALUES ($1, $2, $3, $4);`,
	)

	if err != nil {
		return nil, fmt.Errorf("unable to prepare statement: %v", err)
	}

	defer stmt.Close()

	var ids []uint64 // inserted record identifiers

	for i := 0; i < len(entities); i++ {
		f, ok := entities[i].(*models.Feedback)
		if !ok {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("unable to rollback transaction: %v", err)
			}
			return nil, errors.New("underlying type must be *models.Feedback")
		}

		var sequenceNumber uint64
		err := tx.QueryRowContext(ctx,
			"SELECT nextval('reaction.feedback_id_seq'::regclass);",
		).Scan(&sequenceNumber)

		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("unable to rollback transaction: %v", err)
			}
			return nil, fmt.Errorf("unable to get next seq number: %v", err)
		}

		_, err = stmt.ExecContext(ctx, sequenceNumber, f.UserId, f.ClassroomId, f.Comment)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return nil, fmt.Errorf("unable to rollback transaction: %v", err)
			}
			return nil, fmt.Errorf("unable to insert a record: %v", err)
		}
		ids = append(ids, uint64(sequenceNumber))
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, fmt.Errorf("unable to rollback transaction: %v", err)
		}
		return nil, fmt.Errorf("unable to commit transaction: %v", err)
	}
	return ids, nil
}

// RemoveEntity removes a feedback from the database
func (r *feedbackRepo) RemoveEntity(ctx context.Context, entityId uint64) error {

	// check if record exists
	var dummy uint64
	err := r.db.QueryRowContext(ctx,
		"SELECT 1 FROM reaction.feedback WHERE id=$1;",
		entityId,
	).Scan(&dummy)

	if err == sql.ErrNoRows {
		return fmt.Errorf("no such feedback: %v", err)
	} else if err != nil {
		return fmt.Errorf("unable to remove feedback: %v", err)
	}

	_, err = r.db.ExecContext(ctx,
		"DELETE FROM reaction.feedback WHERE id=$1;",
		entityId,
	)
	if err != nil {
		return fmt.Errorf("unable to remove feedback: %v", err)
	}
	return nil
}

// RemoveEntity retrieves a feedback from the database
func (r *feedbackRepo) DescribeEntity(ctx context.Context, entityId uint64) (models.Entity, error) {

	fb := &models.Feedback{}
	err := r.db.QueryRowContext(ctx,
		"SELECT id, user_id, classroom_id, comment FROM reaction.feedback WHERE id=$1;",
		entityId,
	).Scan(&fb.Id, &fb.UserId, &fb.ClassroomId, &fb.Comment)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no such feedback: %v", err)
	} else if err != nil {
		return nil, fmt.Errorf("unable to get feedback: %v", err)
	}
	return fb, nil
}

// RemoveEntity retrieves a feedback from the database
func (r *feedbackRepo) ListEntities(ctx context.Context, limit, offset uint64) ([]models.Entity, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, user_id, classroom_id, comment FROM reaction.feedback LIMIT $1 OFFSET $2;",
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to get feedbacks: %v", err)
	}
	defer rows.Close()

	var feedbacks []models.Entity
	for rows.Next() {
		fb := &models.Feedback{}
		err := rows.Scan(&fb.Id, &fb.UserId, &fb.ClassroomId, &fb.Comment)
		if err != nil {
			return nil, fmt.Errorf("unable to get feedbacks: %v", err)
		}
		feedbacks = append(feedbacks, fb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to get feedbacks: %v", err)
	}

	return feedbacks, nil
}
