package utils

import (
	"testing"
)

func assertNilError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error must not be returned")
	}
}

func assertNonNilError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("error must be returned")
	}
}
