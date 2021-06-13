package grpc_service

import (
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
)

type GrpcService struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	proposalRepo repo.Repo
	chunks       int
}

// New returns a new Feedback GRPC server
func New(fRepo repo.Repo, pRepo repo.Repo, chunks int) *GrpcService {
	return &GrpcService{feedbackRepo: fRepo, proposalRepo: pRepo, chunks: chunks}
}
