package utils

import (
	"testing"
)

func assertNilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("error must not be returned")
	}
}

func assertNonNilError(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("error must be returned")
	}
}

// dummy is just type that implements Entity interface
// used for testing
type Dummy struct {
	id     uint64
	userId uint64
}

func (d *Dummy) Id() uint64 {
	return d.id
}

func (d *Dummy) UserId() uint64 {
	return d.userId
}
