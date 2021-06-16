package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/prommetrics"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	proposal_service "github.com/ozoncp/ocp-feedback-api/internal/server/proposal_grpc"
	"github.com/ozoncp/ocp-feedback-api/internal/tracer"
	pr "github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	grpcPort int
	chunks   int

	// postgres
	dbConnString   string
	dbMaxOpenConns int
	dbMaxIdleConns int
	// prometheus
	promAddr string
	// kafka
	brokerList string
)

func init() {
	flag.IntVar(&grpcPort, "port", 10000, "GRPC server port")
	flag.IntVar(&chunks, "chunks", 2, "Number of chunks to split into")
	flag.StringVar(&dbConnString, "db-conn", "postgres://postgres:postgres@localhost/postgres", "Database connection string")
	flag.IntVar(&dbMaxOpenConns, "db-MaxOpenConnections", 15, "Number of total open connections to the database")
	flag.IntVar(&dbMaxIdleConns, "db-MaxIdleConnections", 5, "Number of idle connections in the pool")
	flag.StringVar(&promAddr, "prometheus-address", ":2112", "The address to listen on for HTTP requests.")
	flag.StringVar(&brokerList, "broker-address", "127.0.0.1:29092", "List of KAFKA brokers")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := sqlx.Connect("pgx", dbConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the database")
	}
	defer db.Close()

	db.SetMaxOpenConns(dbMaxOpenConns)
	db.SetMaxIdleConns(dbMaxOpenConns)

	log.Info().Msg("Connected to the database")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, ctx := errgroup.WithContext(ctx)
	defer cancel()

	// create asynchronous KAFKA producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionNone
	config.Producer.Flush.Frequency = time.Second

	asyncProducer, err := sarama.NewAsyncProducer(strings.Split(brokerList, ","), config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start Sarama producer:%v")
	}
	prod := producer.New("proposals", asyncProducer)
	prod.Init(ctx)

	// initialize tracer
	closer := tracer.Init("ocp-proposal-api")
	defer closer.Close()

	// create GRPC service
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
			prod,
			prommetrics.New("proposal"),
			chunks,
		),
	)

	group.Go(func() error {
		log.Info().Msg("Serving grpc requests...")
		return grpcServer.Serve(lis)
	})

	group.Go(func() error {
		log.Info().Msgf("Serving Prometheus metrics at %v", promAddr)
		http.Handle("/metrics", promhttp.Handler())
		return http.ListenAndServe(promAddr, nil)
	})

	if err = group.Wait(); err != nil {
		log.Error().Msgf("Terminated abnormally: %v", err)
	}
	log.Error().Msgf("Terminated normally: %v", err)
}
