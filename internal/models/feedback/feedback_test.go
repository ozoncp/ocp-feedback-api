package feedback

import (
	"fmt"
	"testing"
)

func TestFeedbackCtor(t *testing.T) {

	t.Run("valid ctor", func(t *testing.T) {
		want := &feedback{id: 1, userId: 42, classroomId: 50, comment: "objectName"}
		feedback, err := New(want.id, want.userId, want.classroomId, want.comment)

		assertError(t, err, nil)
		assertFeedback(t, feedback, want)
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := New(0, 42, 50, "comment")

		assertError(t, err, ErrInvalidId)
	})

	t.Run("invalid user id", func(t *testing.T) {
		_, err := New(1, 0, 50, "comment")

		assertError(t, err, ErrInvalidUserId)
	})

	t.Run("invalid classroom id", func(t *testing.T) {
		_, err := New(1, 2, 0, "comment")

		assertError(t, err, ErrInvalidClassroomId)
	})

	t.Run("invalid comment", func(t *testing.T) {
		_, err := New(1, 1, 50, "")

		assertError(t, err, ErrInvalidComment)
	})

	t.Run("valid id", func(t *testing.T) {
		var id uint64 = 24
		feedback, _ := New(id, 42, 50, "comment")

		assertUint64(t, feedback.Id(), id)
	})

	t.Run("valid user id", func(t *testing.T) {
		var userId uint64 = 42
		feedback, _ := New(42, userId, 50, "comment")

		assertUint64(t, feedback.UserId(), userId)
	})

	t.Run("valid classroom id", func(t *testing.T) {
		var classroomId uint64 = 42
		feedback, _ := New(42, 50, classroomId, "comment")

		assertUint64(t, feedback.ClassroomId(), classroomId)
	})

	t.Run("valid comment", func(t *testing.T) {
		comment := "comment"
		feedback, _ := New(42, 42, 50, comment)

		assertString(t, feedback.Comment(), comment)
	})

}

func TestFeedbackSetComment(t *testing.T) {
	t.Run("valid comment", func(t *testing.T) {
		feedback, _ := New(24, 42, 50, "comment")
		updatedComment := "updated_comment"
		err := feedback.UpdateComment(updatedComment)

		assertError(t, err, nil)
		assertString(t, feedback.Comment(), updatedComment)
	})

	t.Run("invalid comment", func(t *testing.T) {
		feedback, _ := New(24, 42, 50, "comment")
		updatedComment := ""
		err := feedback.UpdateComment(updatedComment)

		assertError(t, err, ErrInvalidComment)
	})

}

func TestFeedbackStringConversion(t *testing.T) {
	f := &feedback{1, 0, 50, "objectName"}
	want := fmt.Sprintf(formatString, f.Id(), f.UserId(), f.ClassroomId(), f.Comment())
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

	if got.Id() != want.Id() ||
		got.UserId() != want.UserId() ||
		got.Comment() != want.Comment() {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertUint64(t *testing.T, got, want uint64) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
