package grpc_server

import (
	"context"

	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog/log"
)

type grpcServer struct {
	fb.UnimplementedOcpFeedbackApiServer
}

// New returns a new Feedback GRPC server
func New() *grpcServer {
	return &grpcServer{}
}

// CreateFeedbackV1 ...
func (s *grpcServer) CreateFeedbackV1(ctx context.Context,
	req *fb.CreateFeedbackV1Request) (*fb.CreateFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for CreateFeedbackV1: UserId: %v, ClassroomId: %v, Comment: %v",
		req.UserId, req.ClassroomId, req.Comment)

	return &fb.CreateFeedbackV1Response{FeedbackId: 42}, nil
}

// RemoveFeedbackV1 ...
func (s *grpcServer) RemoveFeedbackV1(ctx context.Context,
	req *fb.RemoveFeedbackV1Request) (*fb.RemoveFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for RemoveFeedbackV1: FeedbackId: %v", req.FeedbackId)
	// TODO return codes.NotFound if requested id is not found
	return &fb.RemoveFeedbackV1Response{}, nil
}

// DescribeFeedbackV1 ...
func (s *grpcServer) DescribeFeedbackV1(ctx context.Context,
	req *fb.DescribeFeedbackV1Request) (*fb.DescribeFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for DescribeFeedbackV1: FeedbackId: %v", req.FeedbackId)
	// TODO return codes.NotFound if requested id is not found
	dummy := &fb.Feedback{FeedbackId: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}
	return &fb.DescribeFeedbackV1Response{Feedback: dummy}, nil
}

// ListFeedbacksV1 ...
func (s *grpcServer) ListFeedbacksV1(ctx context.Context,
	req *fb.ListFeedbacksV1Request) (*fb.ListFeedbacksV1Response, error) {

	log.Info().Msgf("Handle request for ListFeedbacksV1: Limit: %v, Offset: %v", req.Limit, req.Offset)
	// TODO:: return codes.OutOfRange if provided offset is invalid
	dummy := []*fb.Feedback{{FeedbackId: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}}
	return &fb.ListFeedbacksV1Response{Feedbacks: dummy}, nil
}

// CreateProposalV1 ...
func (s *grpcServer) CreateProposalV1(ctx context.Context,
	req *fb.CreateProposalV1Request) (*fb.CreateProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateProposalV1: UserId: %v, LessonId: %v, DocumentId: %v",
		req.UserId, req.LessonId, req.DocumentId)

	return &fb.CreateProposalV1Response{ProposalId: 42}, nil

}

// RemoveProposalV1 ...
func (s *grpcServer) RemoveProposalV1(ctx context.Context,
	req *fb.RemoveProposalV1Request) (*fb.RemoveProposalV1Response, error) {

	log.Info().Msgf("Handle request for RemoveProposalV1: ProposalId: %v", req.ProposalId)
	// TODO return codes.NotFound if requested id is not found
	return &fb.RemoveProposalV1Response{}, nil
}

// DescribeProposalV1 ...
func (s *grpcServer) DescribeProposalV1(ctx context.Context,
	req *fb.DescribeProposalV1Request) (*fb.DescribeProposalV1Response, error) {

	log.Info().Msgf("Handle request for DescribeProposalV1: ProposalId: %v", req.ProposalId)
	// TODO return codes.NotFound if requested id is not found

	dummy := &fb.Proposal{ProposalId: 42, UserId: 100, LessonId: 200, DocumentId: 300}
	return &fb.DescribeProposalV1Response{Proposal: dummy}, nil
}

// ListProposalsV1 ...
func (s *grpcServer) ListProposalsV1(ctx context.Context,
	req *fb.ListProposalsV1Request) (*fb.ListProposalsV1Response, error) {

	log.Info().Msgf("Handle request for ListProposalsV1: Limit: %v, Offset: %v", req.Limit, req.Offset)

	// TODO:: return codes.OutOfRange if provided offset is invalid
	dummy := []*fb.Proposal{{ProposalId: 42, UserId: 100, LessonId: 200, DocumentId: 300}}
	return &fb.ListProposalsV1Response{Proposals: dummy}, nil
}
