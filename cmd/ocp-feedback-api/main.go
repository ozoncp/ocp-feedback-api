package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	grpc_server "github.com/ozoncp/ocp-feedback-api/internal/server"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var grpcPort int

func init() {
	flag.IntVar(&grpcPort, "port", 10000, "GRPC server port")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	grpcEndpoint := fmt.Sprintf("localhost:%d", grpcPort)

	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
	}
	log.Info().Msgf("Starting server at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()
	fb.RegisterOcpFeedbackApiServer(grpcServer, grpc_server.New(&repo.InMemoryFeedbackRepo{}, &repo.InMemoryProposalRepo{}))

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}

}
