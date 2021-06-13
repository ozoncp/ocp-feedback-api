package grpc_service

import (
	"context"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateFeedbackV1 saves a new feedback
func (s *GrpcService) CreateFeedbackV1(
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
func (s *GrpcService) CreateMultiFeedbackV1(
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

	chunks, err := utils.SplitSlice(entities, len(entities)/s.chunks) // magic number for now
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &fb.CreateMultiFeedbackV1Response{}

	for i := 0; i < len(chunks); i++ {
		ids, err := s.feedbackRepo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			return res, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.FeedbackIds = append(res.FeedbackIds, ids...)
	}
	return res, nil

}

// RemoveFeedbackV1 removes a feedback
func (s *GrpcService) RemoveFeedbackV1(
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
func (s *GrpcService) DescribeFeedbackV1(
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
func (s *GrpcService) ListFeedbacksV1(
	ctx context.Context,
	req *fb.ListFeedbacksV1Request,
) (*fb.ListFeedbacksV1Response, error) {

	log.Info().Msgf("Handle request for ListFeedbacksV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entities, err := s.feedbackRepo.ListEntities(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "unable to list feedbacks: %v", err)
	}
	var feedbacks []*fb.Feedback

	for i := 0; i < len(entities); i++ {
		f := entities[i].(*models.Feedback)
		feedbacks = append(feedbacks, &fb.Feedback{
			FeedbackId:  f.Id,
			UserId:      f.UserId,
			ClassroomId: f.ClassroomId,
			Comment:     f.Comment,
		})
	}
	return &fb.ListFeedbacksV1Response{Feedbacks: feedbacks}, nil
}
