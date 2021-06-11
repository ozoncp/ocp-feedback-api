package grpc_server

import (
	"context"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	proposalRepo repo.Repo
}

// New returns a new Feedback GRPC server
func New(fRepo repo.Repo, pRepo repo.Repo) *grpcServer {
	return &grpcServer{feedbackRepo: fRepo, proposalRepo: pRepo}
}

// CreateFeedbackV1 saves a new feedback
func (s *grpcServer) CreateFeedbackV1(
	ctx context.Context,
	req *fb.CreateFeedbackV1Request,
) (*fb.CreateFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for CreateFeedbackV1: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	f := &models.Feedback{
		UserId:      req.NewFeedback.UserId,
		ClassroomId: req.NewFeedback.ClassroomId,
		Comment:     req.NewFeedback.Comment,
	}
	ids, err := s.feedbackRepo.AddEntities(ctx, f)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insertion failed: %v", err)
	}
	return &fb.CreateFeedbackV1Response{FeedbackId: ids[0]}, nil
}

// CreateMultiFeedbackV1 creates multiple feedbacks
func (s *grpcServer) CreateMultiFeedbackV1(
	ctx context.Context,
	req *fb.CreateMultiFeedbackV1Request,
) (*fb.CreateMultiFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for CreateMultiFeedbackV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	var entities []models.Entity

	for i := 0; i < len(req.NewFeedbacks); i++ {
		entities = append(entities, &models.Feedback{
			UserId:      req.NewFeedbacks[i].UserId,
			ClassroomId: req.NewFeedbacks[i].ClassroomId,
			Comment:     req.NewFeedbacks[i].Comment,
		})
	}

	chunks, err := utils.SplitSlice(entities, len(entities)/4) // magic number for now
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &fb.CreateMultiFeedbackV1Response{}

	for i := 0; i < len(chunks); i++ {
		ids, err := s.feedbackRepo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.FeedbackIds = append(res.FeedbackIds, ids...)
	}
	return res, nil

}

// RemoveFeedbackV1 removes a feedback
func (s *grpcServer) RemoveFeedbackV1(
	ctx context.Context,
	req *fb.RemoveFeedbackV1Request,
) (*fb.RemoveFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for RemoveFeedbackV1 %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.feedbackRepo.RemoveEntity(ctx, req.FeedbackId); err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete a feedback: %v", err)
	}
	return &fb.RemoveFeedbackV1Response{}, nil
}

// DescribeFeedbackV1 returns a feedback
func (s *grpcServer) DescribeFeedbackV1(
	ctx context.Context,
	req *fb.DescribeFeedbackV1Request,
) (*fb.DescribeFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for DescribeFeedbackV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	entity, err := s.feedbackRepo.DescribeEntity(ctx, req.FeedbackId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to describe a feedback: %v", err)
	}
	f := entity.(*models.Feedback)
	respFeedback := fb.Feedback{
		FeedbackId:  f.Id,
		UserId:      f.UserId,
		ClassroomId: f.ClassroomId,
		Comment:     f.Comment,
	}
	return &fb.DescribeFeedbackV1Response{Feedback: &respFeedback}, nil
}

// ListFeedbacksV1 returns a list of at most 'limit' feedbacks starting from 'offset'
func (s *grpcServer) ListFeedbacksV1(
	ctx context.Context,
	req *fb.ListFeedbacksV1Request,
) (*fb.ListFeedbacksV1Response, error) {

	log.Info().Msgf("Handle request for ListFeedbacksV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO:: return codes.OutOfRange if provided offset is invalid
	dummy := []*fb.Feedback{{FeedbackId: 42, UserId: 100, ClassroomId: 200, Comment: "just_a_comment"}}
	return &fb.ListFeedbacksV1Response{Feedbacks: dummy}, nil
}

// CreateProposalV1 ...
func (s *grpcServer) CreateProposalV1(ctx context.Context,
	req *fb.CreateProposalV1Request) (*fb.CreateProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateProposalV1: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &fb.CreateProposalV1Response{ProposalId: 42}, nil

}

// RemoveProposalV1 ...
func (s *grpcServer) RemoveProposalV1(ctx context.Context,
	req *fb.RemoveProposalV1Request) (*fb.RemoveProposalV1Response, error) {

	log.Info().Msgf("Handle request for RemoveProposalV1: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO return codes.NotFound if requested id is not found
	return &fb.RemoveProposalV1Response{}, nil
}

// DescribeProposalV1 ...
func (s *grpcServer) DescribeProposalV1(ctx context.Context,
	req *fb.DescribeProposalV1Request) (*fb.DescribeProposalV1Response, error) {

	log.Info().Msgf("Handle request for DescribeProposalV1: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO return codes.NotFound if requested id is not found
	dummy := &fb.Proposal{ProposalId: 42, UserId: 100, LessonId: 200, DocumentId: 300}
	return &fb.DescribeProposalV1Response{Proposal: dummy}, nil
}

// ListProposalsV1 ...
func (s *grpcServer) ListProposalsV1(ctx context.Context,
	req *fb.ListProposalsV1Request) (*fb.ListProposalsV1Response, error) {

	log.Info().Msgf("Handle request for ListProposalsV1: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO:: return codes.OutOfRange if provided offset is invalid
	dummy := []*fb.Proposal{{ProposalId: 42, UserId: 100, LessonId: 200, DocumentId: 300}}
	return &fb.ListProposalsV1Response{Proposals: dummy}, nil
}
