package utils

import (
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models/feedback"
	"github.com/ozoncp/ocp-feedback-api/internal/models/proposal"
)

// ConvertFeedback converts slice of feedbacks into a map where key is a feedback id
// and value is a feedback
// If passed slice is nil, error will be returned
// If passed slice containts two equal ids, panic will be invoked
func ConvertFeedback(fbSlice []feedback.Feedback) (map[uint64]feedback.Feedback, error) {
	if fbSlice == nil {
		return nil, errors.New("cannot convert nil slice")
	}

	fbMap := make(map[uint64]feedback.Feedback, len(fbSlice))

	for i := 0; i < len(fbSlice); i++ {
		fb := fbSlice[i]
		if _, present := fbMap[fb.Id()]; present {
			panic("feedback id collision")
		}
		fbMap[fb.Id()] = fb
	}
	return fbMap, nil
}

// ConvertProposal converts slice of proposals into a map where key is a proposal id
// and value is a proposal
// If passed slice is nil, error will be returned
// If passed slice containts two equal ids, panic will be invoked
func ConvertProposal(prSlice []proposal.Proposal) (map[uint64]proposal.Proposal, error) {
	if prSlice == nil {
		return nil, errors.New("cannot convert nil slice")
	}

	prMap := make(map[uint64]proposal.Proposal, len(prSlice))

	for i := 0; i < len(prSlice); i++ {
		pr := prSlice[i]
		if _, present := prMap[pr.Id()]; present {
			panic("proposal id collision")
		}
		prMap[pr.Id()] = pr
	}
	return prMap, nil
}
