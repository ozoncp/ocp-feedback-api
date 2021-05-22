package proposal

import (
	"fmt"
	"testing"
)

func TestProposalCtor(t *testing.T) {

	t.Run("valid object", func(t *testing.T) {
		want := &proposal{id: 1, userId: 2, lessonId: 3, documentId: 4}
		got, err := New(want.Id(), want.UserId(), want.LessonId(), want.DocumentId())

		assertError(t, err, nil)
		assertProposal(t, got, want)
	})

	t.Run("invalid id", func(t *testing.T) {
		p := &proposal{id: 0, userId: 2, lessonId: 3, documentId: 4}
		_, err := New(p.Id(), p.UserId(), p.LessonId(), p.DocumentId())

		assertError(t, err, ErrInvalidId)
	})

	t.Run("invalid user id", func(t *testing.T) {
		p := &proposal{id: 1, userId: 0, lessonId: 3, documentId: 4}
		_, err := New(p.Id(), p.UserId(), p.LessonId(), p.DocumentId())

		assertError(t, err, ErrInvalidUserId)
	})

	t.Run("invalid lesson id", func(t *testing.T) {
		p := &proposal{id: 1, userId: 2, lessonId: 0, documentId: 4}
		_, err := New(p.Id(), p.UserId(), p.LessonId(), p.DocumentId())

		assertError(t, err, ErrInvalidLessonId)
	})

	t.Run("invalid document id", func(t *testing.T) {
		p := &proposal{id: 1, userId: 2, lessonId: 3, documentId: 0}
		_, err := New(p.Id(), p.UserId(), p.LessonId(), p.DocumentId())

		assertError(t, err, ErrInvalidDocumentId)
	})
}

func TestProposalStringConversion(t *testing.T) {
	p := &proposal{id: 1, userId: 2, lessonId: 3, documentId: 4}
	want := fmt.Sprintf(formatString, p.id, p.userId, p.lessonId, p.documentId)
	got := fmt.Sprint(p)

	assertString(t, got, want)
}

func TestProposalErrConversion(t *testing.T) {
	want := "some error"
	err := ProposalErr(want)
	got := err.Error()

	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertProposal(t *testing.T, got, want Proposal) {
	t.Helper()

	if got.Id() != want.Id() ||
		got.UserId() != want.UserId() ||
		got.LessonId() != want.LessonId() ||
		got.DocumentId() != want.DocumentId() {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
