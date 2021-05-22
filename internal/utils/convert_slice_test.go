package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/models/feedback"
	"github.com/ozoncp/ocp-feedback-api/internal/models/proposal"
)

func TestFeedbackConversion(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var fb []feedback.Feedback
		_, err := ConvertFeedback(fb)

		assertNonNilError(t, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		fb := []feedback.Feedback{}
		got, err := ConvertFeedback(fb)
		want := make(map[uint64]feedback.Feedback)

		assertNilError(t, err)
		assertFeedbackMap(t, got, want)
	})

	t.Run("unique ids", func(t *testing.T) {
		fb1, _ := feedback.New(1, 1, "comment1")
		fb2, _ := feedback.New(2, 2, "comment2")
		fbSlice := []feedback.Feedback{fb1, fb2}
		want := map[uint64]feedback.Feedback{1: fb1, 2: fb2}
		got, err := ConvertFeedback(fbSlice)
		assertNilError(t, err)
		assertFeedbackMap(t, got, want)
	})

	t.Run("duplicate ids", func(t *testing.T) {
		fb1, _ := feedback.New(1, 2, "comment1")
		fb2, _ := feedback.New(1, 3, "comment2")
		fbSlice := []feedback.Feedback{fb1, fb2}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("goroutine must enter panic state")
			}
		}()
		_, _ = ConvertFeedback(fbSlice)
		t.Error("goroutine must enter panic state")
	})
}

func TestProposalConversion(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var p []proposal.Proposal
		_, err := ConvertProposal(p)

		assertNonNilError(t, err)
	})

	t.Run("empty slice", func(t *testing.T) {
		p := []proposal.Proposal{}
		got, err := ConvertProposal(p)
		want := make(map[uint64]proposal.Proposal)

		assertNilError(t, err)
		assertProposalMap(t, got, want)
	})

	t.Run("unique ids", func(t *testing.T) {
		p1, _ := proposal.New(1, 2, 3, 4)
		p2, _ := proposal.New(2, 20, 30, 40)
		pSlice := []proposal.Proposal{p1, p2}
		want := map[uint64]proposal.Proposal{1: p1, 2: p2}
		got, err := ConvertProposal(pSlice)
		assertNilError(t, err)
		assertProposalMap(t, got, want)
	})

	t.Run("duplicate ids", func(t *testing.T) {
		p1, _ := proposal.New(1, 2, 3, 4)
		p2, _ := proposal.New(1, 20, 30, 40)
		pSlice := []proposal.Proposal{p1, p2}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("goroutine must enter panic state")
			}
		}()
		_, _ = ConvertProposal(pSlice)
		t.Error("goroutine must enter panic state")
	})
}

func assertFeedbackMap(t *testing.T, got, want map[uint64]feedback.Feedback) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertProposalMap(t *testing.T, got, want map[uint64]proposal.Proposal) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
