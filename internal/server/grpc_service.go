package grpc_service

import (
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
)

const Chunks = 2

type GrpcService struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	proposalRepo repo.Repo
}

// New returns a new Feedback GRPC server
func New(fRepo repo.Repo, pRepo repo.Repo) *GrpcService {
	return &GrpcService{feedbackRepo: fRepo, proposalRepo: pRepo}
}
