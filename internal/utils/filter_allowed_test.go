package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

func TestFilterBlocked(t *testing.T) {
	t.Run("list slice is nil", func(t *testing.T) {
		var list []entity.Entity
		blockList := []uint64{}
		defer assertPanic(t)
		_, _ = FilterAllowed(list, blockList)
		t.Error("goroutine must enter panic state")
	})

	t.Run("blockList slice is nil", func(t *testing.T) {
		list := []entity.Entity{}
		var blockList []uint64
		defer assertPanic(t)
		_, _ = FilterAllowed(list, blockList)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty blockList", func(t *testing.T) {
		list := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
		}
		blockList := []uint64{}

		want := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("empty list", func(t *testing.T) {
		list := []entity.Entity{}
		blockList := []uint64{1, 2, 3, 4}
		want := []entity.Entity{}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("no items filtered", func(t *testing.T) {
		list := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 2, userId: 4},
		}
		blockList := []uint64{3, 4, 5, 6, 7}
		want := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 2, userId: 4},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("unique items filtered", func(t *testing.T) {
		list := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 2, userId: 4},
			&Dummy{id: 3, userId: 5},
		}

		blockList := []uint64{3, 1, 6}
		want := []entity.Entity{
			&Dummy{id: 2, userId: 4},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})

	t.Run("duplicate items filtered", func(t *testing.T) {
		list := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 1, userId: 7},
			&Dummy{id: 2, userId: 4},
			&Dummy{id: 3, userId: 5},
		}

		blockList := []uint64{1, 6}
		want := []entity.Entity{
			&Dummy{id: 2, userId: 4},
			&Dummy{id: 3, userId: 5},
		}

		got, err := FilterAllowed(list, blockList)

		assertNilError(t, err)
		assertEntity(t, got, want)
	})
}

func assertEntity(t *testing.T, got, want []entity.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
