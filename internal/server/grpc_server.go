package grpc_server

import (
	"context"

	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
)

type grpcServer struct {
	fb.UnimplementedOcpFeedbackApiServer
}

// New ...
func New() *grpcServer {
	return &grpcServer{}
}

// CreateFeedbackV1 ...
func (s *grpcServer) CreateFeedbackV1(ctx context.Context, req *fb.CreateFeedbackV1Request) (*fb.CreateFeedbackV1Response, error) {
	return &fb.CreateFeedbackV1Response{FeedbackId: 42}, nil
}

// RemoveFeedbackV1 ...
func (s *grpcServer) RemoveFeedbackV1(ctx context.Context, req *fb.RemoveFeedbackV1Request) (*fb.RemoveFeedbackV1Response, error) {
	// TODO return codes.NotFound if requested id is not found
	return &fb.RemoveFeedbackV1Response{}, nil
}

// DescribeFeedbackV1 ...
func (s *grpcServer) DescribeFeedbackV1(ctx context.Context, req *fb.DescribeFeedbackV1Request) (*fb.DescribeFeedbackV1Response, error) {
	// TODO return codes.NotFound if requested id is not found
	dummy := &fb.Feedback{FeedbackId: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}
	return &fb.DescribeFeedbackV1Response{Feedback: dummy}, nil
}

// ListFeedbacksV1 ...
func (s *grpcServer) ListFeedbacksV1(ctx context.Context, req *fb.ListFeedbacksV1Request) (*fb.ListFeedbacksV1Response, error) {
	// TODO:: return codes.OutOfRange if provided offset is invalid
	dummy := []*fb.Feedback{{FeedbackId: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}}
	return &fb.ListFeedbacksV1Response{Feedbacks: dummy}, nil
}
