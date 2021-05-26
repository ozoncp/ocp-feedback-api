package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

func TestSplitSlice(t *testing.T) {

	t.Run("nil slice", func(t *testing.T) {
		var slice []models.Entity
		defer assertPanic(t)
		_, _ = SplitSlice(slice, 1)
		t.Error("goroutine must enter panic state")
	})

	t.Run("negative chunk size", func(t *testing.T) {
		slice := []models.Entity{}
		defer assertPanic(t)
		_, _ = SplitSlice(slice, -1)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty slice", func(t *testing.T) {
		got, err := SplitSlice([]models.Entity{}, 1)
		want := [][]models.Entity{}

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("zero chunk size", func(t *testing.T) {

		slice := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
		}

		want := [][]models.Entity{
			{&Dummy{id: 1}, &Dummy{id: 3}},
		}

		got, err := SplitSlice(slice, 0)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("slice size divisible by chunk size", func(t *testing.T) {
		slice := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
			&Dummy{id: 5},
			&Dummy{id: 7},
		}

		want := [][]models.Entity{
			{&Dummy{id: 1}, &Dummy{id: 3}},
			{&Dummy{id: 5}, &Dummy{id: 7}},
		}
		got, err := SplitSlice(slice, 2)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("slice size is not divisible by chunk size", func(t *testing.T) {
		slice := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
			&Dummy{id: 5},
			&Dummy{id: 7},
			&Dummy{id: 9},
		}

		want := [][]models.Entity{
			{&Dummy{id: 1}, &Dummy{id: 3}, &Dummy{id: 5}},
			{&Dummy{id: 7}, &Dummy{id: 9}},
		}
		got, err := SplitSlice(slice, 3)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})

	t.Run("chunk size greater than slice size", func(t *testing.T) {
		slice := []models.Entity{
			&Dummy{id: 1},
			&Dummy{id: 3},
		}

		want := [][]models.Entity{
			{&Dummy{id: 1}, &Dummy{id: 3}},
		}
		got, err := SplitSlice(slice, 10)

		assertNilError(t, err)
		assertEntityMatrix(t, got, want)
	})
}

func assertEntityMatrix(t *testing.T, got, want [][]models.Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
