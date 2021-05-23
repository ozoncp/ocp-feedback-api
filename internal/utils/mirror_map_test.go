package utils

import (
	"reflect"
	"testing"
)

func TestMirrorMap(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var dict map[interface{}]interface{}
		defer assertPanic(t)
		_, _ = MirrorMap(dict)
		t.Error("goroutine must enter panic state")
	})

	t.Run("empty map", func(t *testing.T) {
		dict := make(map[interface{}]interface{})
		want := make(map[interface{}]interface{})
		got, err := MirrorMap(dict)

		assertNilError(t, err)
		assertMap(t, got, want)
	})

	t.Run("valid map mirrored", func(t *testing.T) {
		dict := map[interface{}]interface{}{1: "10", 2: "20", 3: "30"}
		want := map[interface{}]interface{}{"10": 1, "20": 2, "30": 3}
		got, err := MirrorMap(dict)

		assertNilError(t, err)
		assertMap(t, got, want)
	})

	t.Run("key collision panic", func(t *testing.T) {
		dict := map[interface{}]interface{}{1: "10", 2: "10", 3: "30"}
		assertPanic(t)
		_, _ = MirrorMap(dict)
		t.Error("goroutine must enter panic state")
	})
}

func assertMap(t *testing.T, got, want map[interface{}]interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
