package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	proposal_service "github.com/ozoncp/ocp-feedback-api/internal/server/proposal_grpc"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	grpcPort int
	chunks   int

	// postgres
	dbUserName      string
	dbPassword      string
	dbHost          string
	dbPort          string
	dbName          string
	dbMaxOpenConns  int
	dbMaxIdleConnts int
)

func init() {
	flag.IntVar(&grpcPort, "port", 10000, "GRPC server port")
	flag.IntVar(&chunks, "chunks", 2, "Number of chunks to split into")
	flag.StringVar(&dbUserName, "db_user", "postgres", "Database user")
	flag.StringVar(&dbPassword, "db_password", "mydummypassword", "Database password")
	flag.StringVar(&dbHost, "db_host", "localhost", "Database address")
	flag.StringVar(&dbPort, "db_port", "5432", "Database port")
	flag.StringVar(&dbName, "db_name", "postgres", "Database name")
	flag.IntVar(&dbMaxOpenConns, "db_MaxOpenConnections", 15, "Number of total open connections to the database")
	flag.IntVar(&dbMaxIdleConnts, "db_MaxIdleConnections", 5, "Number of idle connections in the pool")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	connString := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
	log.Info().Msgf("Connecting to postgres %v...", connString)

	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the database")
	}
	defer db.Close()

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxOpenConns)
	log.Info().Msg("Connected to postgres ...")

	grpcEndpoint := fmt.Sprintf("localhost:%d", grpcPort)
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
	}
	log.Info().Msgf("Starting proposal service at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()

	pr.RegisterOcpProposalApiServer(grpcServer,
		proposal_service.New(
			repo.NewProposalRepo(db),
			chunks,
		),
	)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}
}