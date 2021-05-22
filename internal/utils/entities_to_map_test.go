package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

func TestFeedbackConversion(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var entities []entity.Entity
		_, err := EntitiesToMap(entities)

		assertNonNilError(t, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		entities := []entity.Entity{}
		got, err := EntitiesToMap(entities)
		want := make(map[uint64]entity.Entity)

		assertNilError(t, err)
		assertEntitiesMap(t, got, want)
	})

	t.Run("unique ids", func(t *testing.T) {
		entities := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 2, userId: 7},
			&Dummy{id: 3, userId: 4},
		}
		want := map[uint64]entity.Entity{1: entities[0], 2: entities[1], 3: entities[2]}
		got, err := EntitiesToMap(entities)
		assertNilError(t, err)
		assertEntitiesMap(t, got, want)
	})

	t.Run("duplicate ids", func(t *testing.T) {
		entities := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 1, userId: 7},
			&Dummy{id: 3, userId: 4},
		}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("goroutine must enter panic state")
			}
		}()
		_, _ = EntitiesToMap(entities)
		t.Error("goroutine must enter panic state")
	})
}

func assertEntitiesMap(t *testing.T, got, want map[uint64]entity.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
