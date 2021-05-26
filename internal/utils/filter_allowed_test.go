package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

func TestFilterBlocked(t *testing.T) {
	t.Run("list slice is nil", func(t *testing.T) {
		var list []models.Entity
		blockList := []uint64{}
		defer assertPanic(t)
		_, _ = FilterAllowed(list, blockList)
		t.Error("goroutine must enter panic state")
	})

	t.Run("blockList slice is nil", func(t *testing.T) {
		list := []models.Entity{}
		var blockList []uint64
		defer assertPanic(t)
		_, _ = FilterAllowed(list, blockList)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty blockList", func(t *testing.T) {
		list := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
		}
		blockList := []uint64{}

		want := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("empty list", func(t *testing.T) {
		list := []models.Entity{}
		blockList := []uint64{1, 2, 3, 4}
		want := []models.Entity{}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("no items filtered", func(t *testing.T) {
		list := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 2},
		}
		blockList := []uint64{3, 4, 5, 6, 7}
		want := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 2},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("unique items filtered", func(t *testing.T) {
		list := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 2},
			&Dummy{id: 3},
		}

		blockList := []uint64{3, 1, 6}
		want := []models.Entity{
			&Dummy{id: 2},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("duplicate items filtered", func(t *testing.T) {
		list := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 1},
			&Dummy{id: 2},
			&Dummy{id: 3},
		}

		blockList := []uint64{1, 6}
		want := []models.Entity{
			&Dummy{id: 2},
			&Dummy{id: 3},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})
}

func assertEntity(t *testing.T, got, want []models.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
