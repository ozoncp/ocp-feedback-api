package grpc_service

import (
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
)

const Chunks = 4

type grpcService struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	proposalRepo repo.Repo
}

// New returns a new Feedback GRPC server
func New(fRepo repo.Repo, pRepo repo.Repo) *grpcService {
	return &grpcService{feedbackRepo: fRepo, proposalRepo: pRepo}
}
