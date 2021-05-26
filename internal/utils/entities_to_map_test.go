package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

func TestFeedbackConversion(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var entities []models.Entity
		defer assertPanic(t)
		_, _ = EntitiesToMap(entities)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty slice", func(t *testing.T) {
		entities := []models.Entity{}
		got, err := EntitiesToMap(entities)
		want := make(map[uint64]models.Entity)

		assertNilError(t, err)
		assertEntitiesMap(t, got, want)
	})

	t.Run("unique ids", func(t *testing.T) {
		entities := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 2},
			&Dummy{id: 3},
		}
		want := map[uint64]models.Entity{1: entities[0], 2: entities[1], 3: entities[2]}
		got, err := EntitiesToMap(entities)
		assertNilError(t, err)
		assertEntitiesMap(t, got, want)
	})

	t.Run("duplicate ids", func(t *testing.T) {
		entities := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 1},
			&Dummy{id: 3},
		}
		defer assertPanic(t)
		_, _ = EntitiesToMap(entities)
		t.Error("goroutine must enter panic state")
	})
}

func assertEntitiesMap(t *testing.T, got, want map[uint64]models.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
