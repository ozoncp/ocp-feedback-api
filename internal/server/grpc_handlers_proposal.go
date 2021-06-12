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

// CreateProposalV1 saves a new proposal
func (s *GrpcService) CreateProposalV1(
	ctx context.Context,
	req *fb.CreateProposalV1Request,
) (*fb.CreateProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateProposalV1Request: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	p := &models.Proposal{
		UserId:     req.NewProposal.UserId,
		LessonId:   req.NewProposal.LessonId,
		DocumentId: req.NewProposal.DocumentId,
	}

	ids, err := s.proposalRepo.AddEntities(ctx, p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insertion failed: %v", err)
	}
	return &fb.CreateProposalV1Response{ProposalId: ids[0]}, nil
}

// CreateMultiProposalV1 creates multiple proposals
func (s *GrpcService) CreateMultiProposalV1(
	ctx context.Context,
	req *fb.CreateMultiProposalV1Request,
) (*fb.CreateMultiProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateMultiProposalV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	var entities []models.Entity

	for i := 0; i < len(req.NewProposals); i++ {
		entities = append(entities, &models.Proposal{
			UserId:     req.NewProposals[i].UserId,
			LessonId:   req.NewProposals[i].LessonId,
			DocumentId: req.NewProposals[i].DocumentId,
		})
	}

	chunks, err := utils.SplitSlice(entities, len(entities)/Chunks) // magic number for now
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &fb.CreateMultiProposalV1Response{}

	for i := 0; i < len(chunks); i++ {
		ids, err := s.proposalRepo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.ProposalIds = append(res.ProposalIds, ids...)
	}
	return res, nil

}

// RemoveProposalV1 removes a proposal
func (s *GrpcService) RemoveProposalV1(
	ctx context.Context,
	req *fb.RemoveProposalV1Request,
) (*fb.RemoveProposalV1Response, error) {

	log.Info().Msgf("Handle request for RemoveProposalV1 %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.proposalRepo.RemoveEntity(ctx, req.ProposalId); err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete a proposal: %v", err)
	}
	return &fb.RemoveProposalV1Response{}, nil
}

// DescribeProposalV1 returns a proposal
func (s *GrpcService) DescribeProposalV1(
	ctx context.Context,
	req *fb.DescribeProposalV1Request,
) (*fb.DescribeProposalV1Response, error) {

	log.Info().Msgf("Handle request for DescribeProposalV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	entity, err := s.proposalRepo.DescribeEntity(ctx, req.ProposalId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to describe a proposal: %v", err)
	}
	p := entity.(*models.Proposal)
	respProposal := fb.Proposal{
		ProposalId: p.Id,
		UserId:     p.UserId,
		LessonId:   p.LessonId,
		DocumentId: p.DocumentId,
	}
	return &fb.DescribeProposalV1Response{Proposal: &respProposal}, nil
}

// ListProposalsV1 returns a list of at most 'limit' proposals starting from 'offset'
func (s *GrpcService) ListProposalsV1(
	ctx context.Context,
	req *fb.ListProposalsV1Request,
) (*fb.ListProposalsV1Response, error) {

	log.Info().Msgf("Handle request for ListProposalsV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entities, err := s.proposalRepo.ListEntities(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "unable to list proposals: %v", err)
	}
	var proposals []*fb.Proposal

	for i := 0; i < len(entities); i++ {
		p := entities[i].(*models.Proposal)
		proposals = append(proposals, &fb.Proposal{
			ProposalId: p.Id,
			UserId:     p.UserId,
			LessonId:   p.LessonId,
			DocumentId: p.DocumentId,
		})
	}
	return &fb.ListProposalsV1Response{Proposals: proposals}, nil
}
