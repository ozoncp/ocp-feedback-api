package grpc_service_test

import (
	"net"
	"testing"

	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	grpc_service "github.com/ozoncp/ocp-feedback-api/internal/server"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func startTestGrpcServer(t *testing.T,
	feedbackRepo repo.Repo,
	proposalRepo repo.Repo,
) string {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	service := grpc_service.New(feedbackRepo, proposalRepo)
	grpcServer := grpc.NewServer()
	fb.RegisterOcpFeedbackApiServer(grpcServer, service)
	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)
	go func() {
		err := grpcServer.Serve(listener)
		require.NoError(t, err)
	}()
	return listener.Addr().String()
}

func newTestGrpcClient(t *testing.T, serverAddress string) fb.OcpFeedbackApiClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return fb.NewOcpFeedbackApiClient(conn)
}
