package proposal_grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	oplog "github.com/opentracing/opentracing-go/log"
	"github.com/ozoncp/ocp-feedback-api/internal/models"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/prommetrics"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	"github.com/ozoncp/ocp-feedback-api/internal/utils"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProposalService struct {
	pr.UnimplementedOcpProposalApiServer
	repo   repo.Repo
	prod   producer.Producer
	prom   prommetrics.PromMetrics
	chunks int
}

// New returns a new Proposal GRPC server
func New(repo repo.Repo,
	prod producer.Producer,
	prom prommetrics.PromMetrics,
	chunks int,
) *ProposalService {
	return &ProposalService{
		repo:   repo,
		chunks: chunks,
		prom:   prom,
		prod:   prod,
	}
}

// CreateProposalV1 saves a new proposal
func (s *ProposalService) CreateProposalV1(
	ctx context.Context,
	req *pr.CreateProposalV1Request,
) (*pr.CreateProposalV1Response, error) {

	log.Info().Msgf("Handle request for CreateProposalV1: %v", req)

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

	ids, err := s.repo.AddEntities(ctx, p)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insertion failed: %v", err)
	}

	s.prod.SendEvent(producer.CreateEvent(producer.Create, ids[0]))
	s.prom.IncCreate()

	return &pr.CreateProposalV1Response{ProposalId: ids[0]}, nil
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

	rootspan, spanctx := opentracing.StartSpanFromContext(ctx, "CreateMultiProposalV1")
	defer rootspan.Finish()

	var entities []models.Entity
	for i := 0; i < len(req.Proposals); i++ {
		entities = append(entities, &models.Proposal{
			UserId:     req.Proposals[i].UserId,
			LessonId:   req.Proposals[i].LessonId,
			DocumentId: req.Proposals[i].DocumentId,
		})
	}

	chunks, err := utils.SplitSlice(entities, s.chunks)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := &pr.CreateMultiProposalV1Response{}

	// try to insert into database one chunk per transaction
	// if transaction fails, only those IDs which have been already added successfully
	// will be returned to the client
	for i := 0; i < len(chunks); i++ {
		span, _ := opentracing.StartSpanFromContext(spanctx, "batch")

		addedIds, err := s.repo.AddEntities(ctx, chunks[i]...)
		if err != nil {
			span.LogFields(oplog.Uint64("batch size", 0))
			span.Finish()
			return res, status.Errorf(codes.Internal, "bulk insertion failed: %v", err)
		}
		res.Proposals = append(res.Proposals, addedIds...)

		span.LogFields(oplog.Uint64("batch size", calculateSize(chunks[i]...)))
		span.Finish()

		for _, id := range addedIds {
			s.prod.SendEvent(producer.CreateEvent(producer.Create, id))
			s.prom.IncCreate()
		}
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
	err := s.repo.RemoveEntity(ctx, req.ProposalId)

	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.prod.SendEvent(producer.CreateEvent(producer.Remove, req.ProposalId))
	s.prom.IncRemove()

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

	entity, err := s.repo.DescribeEntity(ctx, req.ProposalId)

	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	p := entity.(*models.Proposal)
	respProposal := pr.Proposal{
		ProposalId: p.Id,
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

	entities, err := s.repo.ListEntities(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "unable to list proposals: %v", err)
	}

	var proposals []*pr.Proposal
	for i := 0; i < len(entities); i++ {
		p := entities[i].(*models.Proposal)
		proposals = append(proposals, &pr.Proposal{
			ProposalId: p.Id,
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
		Id:         req.Proposal.ProposalId,
		UserId:     req.Proposal.UserId,
		LessonId:   req.Proposal.LessonId,
		DocumentId: req.Proposal.DocumentId,
	}

	err := s.repo.UpdateEntity(ctx, p)

	if err == repo.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	s.prod.SendEvent(producer.CreateEvent(producer.Update, req.Proposal.ProposalId))
	s.prom.IncUpdate()

	return &pr.UpdateProposalV1Response{}, nil
}

func calculateSize(entities ...models.Entity) uint64 {
	var batchSize uint64
	for i := 0; i < len(entities); i++ {
		batchSize += entities[i].(*models.Proposal).Size()
	}
	return batchSize
}
