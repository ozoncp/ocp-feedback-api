package utils

import (
	"reflect"
	"testing"
)

func TestFilterBlocked(t *testing.T) {
	t.Run("list slice is nil", func(t *testing.T) {
		var list []interface{}
		blockList := []interface{}{}
		_, err := FilterAllowed(list, blockList)

		assertNonNilError(t, err)
	})

	t.Run("blockList slice is nil", func(t *testing.T) {
		list := []interface{}{}
		var blockList []interface{}
		_, err := FilterAllowed(list, blockList)

		assertNonNilError(t, err)
	})

	t.Run("empty blockList", func(t *testing.T) {
		list := []interface{}{1, 2, 3}
		blockList := []interface{}{}

		got, err := FilterAllowed(list, blockList)
		want := []interface{}{1, 2, 3}

		assertNilError(t, err)
		assertSlice(t, got, want)
	})

	t.Run("empty list", func(t *testing.T) {
		list := []interface{}{}
		allowList := []interface{}{1, 2, 3}

		got, err := FilterAllowed(list, allowList)
		want := []interface{}{}

		assertNilError(t, err)
		assertSlice(t, got, want)
	})

	t.Run("no items filtered", func(t *testing.T) {
		list := []interface{}{1, 2, 3}
		blockList := []interface{}{4, 5, 6}
		got, err := FilterAllowed(list, blockList)
		want := []interface{}{1, 2, 3}

		assertNilError(t, err)
		assertSlice(t, got, want)
	})

	t.Run("unique items filtered", func(t *testing.T) {
		list := []interface{}{1, 2, 3}
		blockList := []interface{}{3, 1, 6}
		got, err := FilterAllowed(list, blockList)
		want := []interface{}{2}

		assertNilError(t, err)
		assertSlice(t, got, want)
	})

	t.Run("duplicate items filtered", func(t *testing.T) {
		list := []interface{}{1, 2, 3, 3}
		blockList := []interface{}{3, 5, 6}
		got, err := FilterAllowed(list, blockList)
		want := []interface{}{1, 2}

		assertNilError(t, err)
		assertSlice(t, got, want)
	})
}

func assertSlice(t *testing.T, got, want []interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
