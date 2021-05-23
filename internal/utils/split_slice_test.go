package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

func TestSplitSlice(t *testing.T) {

	t.Run("nil slice", func(t *testing.T) {
		var slice []entity.Entity
		defer assertPanic(t)
		_, _ = SplitSlice(slice, 1)
		t.Error("goroutine must enter panic state")
	})

	t.Run("negative chunk size", func(t *testing.T) {
		slice := []entity.Entity{}
		defer assertPanic(t)
		_, _ = SplitSlice(slice, -1)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty slice", func(t *testing.T) {
		got, err := SplitSlice([]entity.Entity{}, 1)
		want := [][]entity.Entity{}

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("zero chunk size", func(t *testing.T) {

		slice := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
		}

		want := [][]entity.Entity{
			{&Dummy{id: 1, userId: 2}, &Dummy{id: 3, userId: 4}},
		}

		got, err := SplitSlice(slice, 0)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("slice size divisible by chunk size", func(t *testing.T) {
		slice := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
			&Dummy{id: 5, userId: 6},
			&Dummy{id: 7, userId: 8},
		}

		want := [][]entity.Entity{
			{&Dummy{id: 1, userId: 2}, &Dummy{id: 3, userId: 4}},
			{&Dummy{id: 5, userId: 6}, &Dummy{id: 7, userId: 8}},
		}
		got, err := SplitSlice(slice, 2)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("slice size is not divisible by chunk size", func(t *testing.T) {
		slice := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
			&Dummy{id: 5, userId: 6},
			&Dummy{id: 7, userId: 8},
			&Dummy{id: 9, userId: 10},
		}

		want := [][]entity.Entity{
			{&Dummy{id: 1, userId: 2}, &Dummy{id: 3, userId: 4}, &Dummy{id: 5, userId: 6}},
			{&Dummy{id: 7, userId: 8}, &Dummy{id: 9, userId: 10}},
		}
		got, err := SplitSlice(slice, 3)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("chunk size greater than slice size", func(t *testing.T) {
		slice := []entity.Entity{
			&Dummy{id: 1, userId: 2},
			&Dummy{id: 3, userId: 4},
		}

		want := [][]entity.Entity{
			{&Dummy{id: 1, userId: 2}, &Dummy{id: 3, userId: 4}},
		}
		got, err := SplitSlice(slice, 10)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})
}

func assertEntityMatrix(t *testing.T, got, want [][]entity.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
