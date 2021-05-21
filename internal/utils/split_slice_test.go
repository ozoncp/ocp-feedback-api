package utils

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {

	t.Run("nil slice", func(t *testing.T) {
		var slice []interface{}
		_, err := SplitSlice(slice, 1)

		assertNonNilError(t, err)
	})

	t.Run("negative chunk size", func(t *testing.T) {
		slice := []interface{}{}
		_, err := SplitSlice(slice, -1)

		assertNonNilError(t, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []interface{}{}
		got, err := SplitSlice(slice, 1)
		want := [][]interface{}{}

		assertNilError(t, err)
		assert2DSlice(t, got, want)
	})

	t.Run("zero chunk size", func(t *testing.T) {
		slice := []interface{}{1, 2, 3}
		got, err := SplitSlice(slice, 0)
		want := [][]interface{}{{1, 2, 3}}

		assertNilError(t, err)
		assert2DSlice(t, got, want)
	})

	t.Run("slice size divisible by chunk size", func(t *testing.T) {
		slice := []interface{}{1, 2, 3, 4, 5, 6}
		got, err := SplitSlice(slice, 2)
		want := [][]interface{}{{1, 2}, {3, 4}, {5, 6}}

		assertNilError(t, err)
		assert2DSlice(t, got, want)
	})

	t.Run("slice size is not divisible by chunk size", func(t *testing.T) {
		slice := []interface{}{1, 2, 3, 4, 5, 6}
		got, err := SplitSlice(slice, 4)
		want := [][]interface{}{{1, 2, 3, 4}, {5, 6}}

		assertNilError(t, err)
		assert2DSlice(t, got, want)
	})

	t.Run("chunk size greater than slice size", func(t *testing.T) {
		slice := []interface{}{1, 2, 3}
		got, err := SplitSlice(slice, 10)
		want := [][]interface{}{{1, 2, 3}}

		assertNilError(t, err)
		assert2DSlice(t, got, want)
	})
}

func assert2DSlice(t *testing.T, got, want [][]interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
