package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	flag.Parse()
	cfg, err := cfg.Read(configName, configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to read config file")
	}

	db := createDatabase(cfg)
	defer db.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var group errgroup.Group
	ctx, cancel := context.WithCancel(context.Background())

	prod := createKafkaProducer(ctx, cfg)

	// initialize tracer
	closer := tracer.Init("ocp-feedback-api")
	defer closer.Close()

	lis, grpcServer := createGRPCService(cfg, db, prod)
	metricsServer := createMetricsServer(cfg)
	gwServer, err := createGateway(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start gateway")
	}

	go func() {
		<-signals
		fmt.Println("kek")
		if err := metricsServer.Shutdown(ctx); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}
		fmt.Println("kek1")
		grpcServer.GracefulStop()
		cancel()
	}()

	group.Go(func() error {
		log.Info().Msg("Serving grpc requests...")
		return grpcServer.Serve(lis)
	})

	group.Go(func() error {
		log.Info().Msg("Serving gateway requests...")
		return gwServer.ListenAndServe()
	})

	group.Go(func() error {
		return metricsServer.ListenAndServe()
	})

	if err = group.Wait(); err != http.ErrServerClosed {
		log.Fatal().Msgf("Terminated abnormally: %v", err)
	}
	log.Info().Msgf("Terminated normally")

}

func createDatabase(cfg *cfg.Config) *sqlx.DB {
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the database")
	}
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	log.Info().Msg("Connected to the database")
	return db
}

func createKafkaProducer(ctx context.Context, cfg *cfg.Config) producer.Producer {
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
	return prod
}

func createGRPCService(cfg *cfg.Config, db *sqlx.DB, prod producer.Producer) (net.Listener, *grpc.Server) {

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
	return lis, grpcServer
}

func createMetricsServer(cfg *cfg.Config) *http.Server {
	mux := http.NewServeMux()
	mux.Handle(cfg.Prometheus.URI, promhttp.Handler())
	addr := fmt.Sprintf("%v:%v", cfg.Prometheus.Host, cfg.Prometheus.Port)
	log.Info().Msgf("Serving Prometheus metrics at %v%v", addr, cfg.Prometheus.URI)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return srv
}

func createGateway(ctx context.Context, cfg *cfg.Config) (*http.Server, error) {
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcEndpoint := fmt.Sprintf("%v:%v", cfg.GRPC.Host, cfg.GRPC.Port)

	if err := fb.RegisterOcpFeedbackApiHandlerFromEndpoint(
		ctx, gwmux, grpcEndpoint, opts,
	); err != nil {
		return nil, err
	}

	mux.Handle("/swagger/", swaggerMiddleware(cfg))
	mux.Handle("/", gwmux)

	addr := fmt.Sprintf("%v:%v", cfg.Gateway.Host, cfg.Gateway.Port)
	log.Info().Msgf("Serving http gateway at %v", addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return srv, nil
}

func swaggerMiddleware(cfg *cfg.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join(cfg.Gateway.Swagger, p)
		http.ServeFile(w, r, p)
	})
}
