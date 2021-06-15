package feedback_grpc

import (
	"context"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FeedbackService struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	prod         producer.Producer
	chunks       int
}

// New returns a new Feedback GRPC service
func New(fRepo repo.Repo, producer producer.Producer, chunks int) *FeedbackService {
	return &FeedbackService{
		feedbackRepo: fRepo,
		prod:         producer,
		chunks:       chunks}
}

// CreateFeedbackV1 saves a new feedback
func (s *FeedbackService) CreateFeedbackV1(
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
		UserId:      req.Feedback.UserId,
		ClassroomId: req.Feedback.ClassroomId,
		Comment:     req.Feedback.Comment,
	}
	ids, err := s.feedbackRepo.AddEntities(ctx, f)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insertion failed: %v", err)
	}
	s.prod.SendEvent(producer.CreateEvent(producer.Create, ids[0]))
	return &fb.CreateFeedbackV1Response{FeedbackId: ids[0]}, nil
}

// CreateMultiFeedbackV1 creates multiple feedbacks
func (s *FeedbackService) CreateMultiFeedbackV1(
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

	for i := 0; i < len(req.Feedbacks); i++ {
		entities = append(entities, &models.Feedback{
			UserId:      req.Feedbacks[i].UserId,
			ClassroomId: req.Feedbacks[i].ClassroomId,
			Comment:     req.Feedbacks[i].Comment,
		})
	}

	chunks, err := utils.SplitSlice(entities, len(entities)/s.chunks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &fb.CreateMultiFeedbackV1Response{}

	// try to insert into database one chunk per transaction
	// if transaction fails, only those IDs which have been already added successfully
	// will be returned to the client
	for i := 0; i < len(chunks); i++ {
		ids, err := s.feedbackRepo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			return res, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.FeedbackId = append(res.FeedbackId, ids...)
		for i := 0; i < len(ids); i++ {
			s.prod.SendEvent(producer.CreateEvent(producer.Create, ids[i]))
		}

	}
	return res, nil

}

// RemoveFeedbackV1 removes a feedback
func (s *FeedbackService) RemoveFeedbackV1(
	ctx context.Context,
	req *fb.RemoveFeedbackV1Request,
) (*fb.RemoveFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for RemoveFeedbackV1 %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := s.feedbackRepo.RemoveEntity(ctx, req.FeedbackId)
	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	s.prod.SendEvent(producer.CreateEvent(producer.Remove, req.FeedbackId))
	return &fb.RemoveFeedbackV1Response{}, nil
}

// DescribeFeedbackV1 returns a feedback
func (s *FeedbackService) DescribeFeedbackV1(
	ctx context.Context,
	req *fb.DescribeFeedbackV1Request,
) (*fb.DescribeFeedbackV1Response, error) {

	log.Info().Msgf("Handle request for DescribeFeedbackV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	entity, err := s.feedbackRepo.DescribeEntity(ctx, req.FeedbackId)
	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
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
func (s *FeedbackService) ListFeedbacksV1(
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

// UpdateFeedbackV1 updates a feedback
func (s *FeedbackService) UpdateFeedbackV1(
	ctx context.Context,
	req *fb.UpdateFeedbackV1Request,
) (*fb.UpdateFeedbackV1Response, error) {
	log.Info().Msgf("Handle request for UpdateFeedbackV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	f := &models.Feedback{
		Id:          req.Feedback.FeedbackId,
		UserId:      req.Feedback.UserId,
		ClassroomId: req.Feedback.ClassroomId,
		Comment:     req.Feedback.Comment,
	}

	err := s.feedbackRepo.UpdateEntity(ctx, f)
	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	s.prod.SendEvent(producer.CreateEvent(producer.Update, req.Feedback.FeedbackId))
	return &fb.UpdateFeedbackV1Response{}, nil
}
