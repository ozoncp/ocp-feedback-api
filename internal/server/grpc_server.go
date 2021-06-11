package grpc_server

import (
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
)

const Chunks = 4

type grpcServer struct {
	fb.UnimplementedOcpFeedbackApiServer
	feedbackRepo repo.Repo
	proposalRepo repo.Repo
}

// New returns a new Feedback GRPC server
func New(fRepo repo.Repo, pRepo repo.Repo) *grpcServer {
	return &grpcServer{feedbackRepo: fRepo, proposalRepo: pRepo}
}
