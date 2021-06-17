package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Shopify/sarama"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	cfg "github.com/ozoncp/ocp-feedback-api/internal/config"
	"github.com/ozoncp/ocp-feedback-api/internal/producer"
	"github.com/ozoncp/ocp-feedback-api/internal/prommetrics"
	"github.com/ozoncp/ocp-feedback-api/internal/repo"
	feedback_service "github.com/ozoncp/ocp-feedback-api/internal/server/feedback_grpc"
	"github.com/ozoncp/ocp-feedback-api/internal/tracer"
	fb "github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	configName string
	configPath string
)

func init() {
	flag.StringVar(&configName, "config_name", "config", "Name of a .yml config file")
	flag.StringVar(&configPath, "config_path", ".", "Config file path")
}

func main() {
	flag.Parse()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cfg, err := cfg.Read(configName, configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to read config file")
	}

	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the database")
	}
	defer db.Close()

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)

	log.Info().Msg("Connected to the database")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, ctx := errgroup.WithContext(ctx)
	defer cancel()

	// create asynchronous KAFKA producer
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.RequiredAcks = sarama.WaitForLocal // Only wait for the leader to ack
	saramaCfg.Producer.Compression = sarama.CompressionNone
	saramaCfg.Producer.Flush.Frequency = time.Second

	asyncProducer, err := sarama.NewAsyncProducer(cfg.Kafka.Brokers, saramaCfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start Sarama producer")
	}
	prod := producer.New("feedbacks", asyncProducer)
	prod.Init(ctx)

	// initialize tracer
	closer := tracer.Init("ocp-feedback-api")
	defer closer.Close()

	// create GRPC service
	grpcEndpoint := fmt.Sprintf("%v:%v", cfg.GRPC.Host, cfg.GRPC.Port)
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
			prommetrics.New("feedback"),
			cfg.General.Chunks,
		),
	)

	group.Go(func() error {
		log.Info().Msg("Serving grpc requests...")
		return grpcServer.Serve(lis)
	})

	group.Go(func() error {
		log.Info().Msgf("Serving Prometheus metrics at %v", cfg.Prometheus.URI)
		http.Handle(cfg.Prometheus.URI, promhttp.Handler())
		addr := fmt.Sprintf("%v:%v", cfg.Prometheus.Host, cfg.Prometheus.Port)
		return http.ListenAndServe(addr, nil)
	})

	if err = group.Wait(); err != nil {
		log.Fatal().Msgf("Terminated abnormally: %v", err)
	}
	log.Info().Msgf("Terminated normally: %v", err)
}
