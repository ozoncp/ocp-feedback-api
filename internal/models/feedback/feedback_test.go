package feedback

import (
	"fmt"
	"testing"
)

func TestFeedbackCtor(t *testing.T) {

	t.Run("valid object", func(t *testing.T) {
		want := &feedback{1, 42, "objectName"}
		feedback, err := New(want.Id(), want.UserId(), want.Comment())

		assertError(t, err, nil)
		assertFeedback(t, feedback, want)
	})

	t.Run("invalid id", func(t *testing.T) {
		feedback := &feedback{0, 1, "objectName"}
		_, err := New(feedback.Id(), feedback.UserId(), feedback.Comment())

		assertError(t, err, ErrInvalidId)
	})

	t.Run("invalid user id", func(t *testing.T) {
		feedback := &feedback{1, 0, "objectName"}
		_, err := New(feedback.Id(), feedback.UserId(), feedback.Comment())

		assertError(t, err, ErrInvalidUserId)
	})

	t.Run("invalid comment", func(t *testing.T) {
		feedback := &feedback{1, 42, ""}
		_, err := New(feedback.Id(), feedback.UserId(), feedback.Comment())

		assertError(t, err, ErrInvalidComment)
	})
}

func TestFeedbackStringConversion(t *testing.T) {
	f := &feedback{1, 0, "objectName"}
	want := fmt.Sprintf(formatString, f.Id(), f.UserId(), f.Comment())
	got := fmt.Sprint(f)

	assertString(t, got, want)
}

func TestFeedbackErrConversion(t *testing.T) {
	want := "some error"
	err := FeedbackErr(want)
	got := err.Error()

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertFeedback(t *testing.T, got, want Feedback) {
	t.Helper()

	if got.UserId() != want.UserId() || got.Comment() != want.Comment() {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
