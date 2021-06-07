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
	//log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &fb.CreateFeedbackV1Response{Id: 42}, nil
}

// RemoveFeedbackV1 ...
func (s *grpcServer) RemoveFeedbackV1(ctx context.Context, req *fb.RemoveFeedbackV1Request) (*fb.RemoveFeedbackV1Response, error) {
	// TODO use grpc status code to indicate id_not_found
	return &fb.RemoveFeedbackV1Response{}, nil
}

// DescribeFeedbackV1 ...
func (s *grpcServer) DescribeFeedbackV1(ctx context.Context, req *fb.DescribeFeedbackV1Request) (*fb.DescribeFeedbackV1Response, error) {
	dummy := &fb.Feedback{Id: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}
	return &fb.DescribeFeedbackV1Response{Feedback: dummy}, nil
}

// ListFeedbacksV1 ...
func (s *grpcServer) ListFeedbacksV1(ctx context.Context, req *fb.ListFeedbacksV1Request) (*fb.ListFeedbacksV1Response, error) {
	dummy := []*fb.Feedback{{Id: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}}
	return &fb.ListFeedbacksV1Response{Feedbacks: dummy}, nil
}
