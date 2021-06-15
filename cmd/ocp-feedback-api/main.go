package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	feedback_service "github.com/ozoncp/ocp-feedback-api/internal/server/feedback_grpc"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var (
	grpcPort int
	chunks   int

	// postgres
	dbUserName     string
	dbPassword     string
	dbHost         string
	dbPort         string
	dbName         string
	dbMaxOpenConns int
	dbMaxIdleConns int
)

func init() {
	flag.IntVar(&grpcPort, "port", 10000, "GRPC server port")
	flag.IntVar(&chunks, "chunks", 2, "Number of chunks to split into")
	flag.StringVar(&dbUserName, "db_user", "postgres", "Database user")
	flag.StringVar(&dbPassword, "db_password", "postgres", "Database password")
	flag.StringVar(&dbHost, "db_host", "localhost", "Database address")
	flag.StringVar(&dbPort, "db_port", "5432", "Database port")
	flag.StringVar(&dbName, "db_name", "postgres", "Database name")
	flag.IntVar(&dbMaxOpenConns, "db_MaxOpenConnections", 15, "Number of total open connections to the database")
	flag.IntVar(&dbMaxIdleConns, "db_MaxIdleConnections", 5, "Number of idle connections in the pool")
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
	db.SetMaxIdleConns(dbMaxIdleConns)
	log.Info().Msg("Connected to postgres")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create asynchronous KAFKA producer
	prod, err := producer.New("feedbacks")
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}
	prod.Init(ctx)

	grpcEndpoint := fmt.Sprintf("localhost:%d", grpcPort)
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot start feedback grpc server at %v", grpcEndpoint)
	}

	log.Info().Msgf("Starting feedback service at %v...", grpcEndpoint)

	grpcServer := grpc.NewServer()

	fb.RegisterOcpFeedbackApiServer(grpcServer,
		feedback_service.New(
			repo.NewFeedbackRepo(db),
			prod,
			chunks,
		),
	)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Cannot accept connections")
	}
}
