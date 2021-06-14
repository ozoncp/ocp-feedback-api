package proposal_grpc

import (
	"context"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProposalService struct {
	pr.UnimplementedOcpProposalApiServer
	proposalRepo repo.Repo
	chunks       int
}

// New returns a new Feedback GRPC server
func New(pRepo repo.Repo, chunks int) *ProposalService {
	return &ProposalService{proposalRepo: pRepo, chunks: chunks}
}

// CreateProposalV1 saves a new proposal
func (s *ProposalService) CreateProposalV1(
	ctx context.Context,
	req *pr.CreateProposalV1Request,
) (*pr.CreateProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateProposalV1Request: %v", req)
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	p := &models.Proposal{
		UserId:     req.Proposal.UserId,
		LessonId:   req.Proposal.LessonId,
		DocumentId: req.Proposal.DocumentId,
	}

	ids, err := s.proposalRepo.AddEntities(ctx, p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insertion failed: %v", err)
	}
	return &pr.CreateProposalV1Response{Proposal: ids[0]}, nil
}

// CreateMultiProposalV1 creates multiple proposals
func (s *ProposalService) CreateMultiProposalV1(
	ctx context.Context,
	req *pr.CreateMultiProposalV1Request,
) (*pr.CreateMultiProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateMultiProposalV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"request is invalid: %v",
			err.Error())
	}

	var entities []models.Entity

	for i := 0; i < len(req.Proposals); i++ {
		entities = append(entities, &models.Proposal{
			UserId:     req.Proposals[i].UserId,
			LessonId:   req.Proposals[i].LessonId,
			DocumentId: req.Proposals[i].DocumentId,
		})
	}

	chunks, err := utils.SplitSlice(entities, len(entities)/s.chunks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pr.CreateMultiProposalV1Response{}

	// try to insert into database one chunk per transaction
	// if transaction fails, only those IDs which have been already added successfully
	// will be returned to the client
	for i := 0; i < len(chunks); i++ {
		ids, err := s.proposalRepo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			return res, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.Proposals = append(res.Proposals, ids...)
	}
	return res, nil

}

// RemoveProposalV1 removes a proposal
func (s *ProposalService) RemoveProposalV1(
	ctx context.Context,
	req *pr.RemoveProposalV1Request,
) (*pr.RemoveProposalV1Response, error) {

	log.Info().Msgf("Handle request for RemoveProposalV1 %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.proposalRepo.RemoveEntity(ctx, req.Proposal); err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to delete a proposal: %v", err)
	}
	return &pr.RemoveProposalV1Response{}, nil
}

// DescribeProposalV1 returns a proposal
func (s *ProposalService) DescribeProposalV1(
	ctx context.Context,
	req *pr.DescribeProposalV1Request,
) (*pr.DescribeProposalV1Response, error) {

	log.Info().Msgf("Handle request for DescribeProposalV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	entity, err := s.proposalRepo.DescribeEntity(ctx, req.Proposal)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to describe a proposal: %v", err)
	}
	p := entity.(*models.Proposal)
	respProposal := pr.Proposal{
		Id:         p.Id,
		UserId:     p.UserId,
		LessonId:   p.LessonId,
		DocumentId: p.DocumentId,
	}
	return &pr.DescribeProposalV1Response{Proposal: &respProposal}, nil
}

// ListProposalsV1 returns a list of at most 'limit' proposals starting from 'offset'
func (s *ProposalService) ListProposalsV1(
	ctx context.Context,
	req *pr.ListProposalsV1Request,
) (*pr.ListProposalsV1Response, error) {

	log.Info().Msgf("Handle request for ListProposalsV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	entities, err := s.proposalRepo.ListEntities(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "unable to list proposals: %v", err)
	}
	var proposals []*pr.Proposal

	for i := 0; i < len(entities); i++ {
		p := entities[i].(*models.Proposal)
		proposals = append(proposals, &pr.Proposal{
			Id:         p.Id,
			UserId:     p.UserId,
			LessonId:   p.LessonId,
			DocumentId: p.DocumentId,
		})
	}
	return &pr.ListProposalsV1Response{Proposals: proposals}, nil
}

// UpdatePropsalV1 updates a proposal
func (s *ProposalService) UpdateProposalV1(
	ctx context.Context,
	req *pr.UpdateProposalV1Request,
) (*pr.UpdateProposalV1Response, error) {
	log.Info().Msgf("Handle request for UpdateProposalV1: %v", req)

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	p := &models.Proposal{
		Id:         req.Proposal.Id,
		UserId:     req.Proposal.UserId,
		LessonId:   req.Proposal.LessonId,
		DocumentId: req.Proposal.DocumentId,
	}

	if err := s.proposalRepo.UpdateEntity(ctx, p); err != nil {
		return nil, status.Errorf(codes.NotFound, "unable to update a proposal: %v", err)
	}
	return &pr.UpdateProposalV1Response{}, nil
}
