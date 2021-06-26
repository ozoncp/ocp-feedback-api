package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type General struct {
	Chunks int
}

type GRPC struct {
	Port int
}

type Gateway struct {
	Port    int
	Swagger string
}

type Postgres struct {
	ConnString   string
	MaxOpenConns int
	MaxIdleConns int
}

type Prometheus struct {
	URI  string
	Port int
}

type Jaeger struct {
	Host string
}

type Kafka struct {
	Brokers []string
}

type Config struct {
	General    General
	GRPC       GRPC
	Gateway    Gateway
	Postgres   Postgres
	Prometheus Prometheus
	Kafka      Kafka
	Jaeger     Jaeger
}

func Read(name, path string) (*Config, error) {

	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := &Config{}

	chunks, err := getEnvAsInt("CHUNKS")
	if err != nil {
		return nil, err
	}
	cfg.General.Chunks = chunks

	cfg.GRPC.Port = viper.GetInt("grpc.port")

	// get postgres credentials
	pgUser, err := getEnvAsString("PGUSER")
	if err != nil {
		return nil, err
	}

	pgHost, err := getEnvAsString("PGHOST")
	if err != nil {
		return nil, err
	}

	pgPort, err := getEnvAsString("PGPORT")
	if err != nil {
		return nil, err
	}

	pgDb, err := getEnvAsString("PGDATABASE")
	if err != nil {
		return nil, err
	}

	pgPsw, err := getEnvAsString("PGPASSWORD")
	if err != nil {
		return nil, err
	}

	cfg.Postgres.ConnString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		pgUser, pgPsw, pgHost, pgPort, pgDb,
	)

	cfg.Postgres.MaxOpenConns = viper.GetInt("postgres.max_open_conns")
	cfg.Postgres.MaxIdleConns = viper.GetInt("postgres.max_idle_conns")

	cfg.Gateway.Port = viper.GetInt("gateway.port")
	cfg.Gateway.Swagger = viper.GetString("gateway.swagger")

	cfg.Kafka.Brokers, err = getEnvAsSlice("KAFKA_BROKERS", ",")
	if err != nil {
		return nil, err
	}
	cfg.Jaeger.Host, err = getEnvAsString("JAEGER_HOST")
	if err != nil {
		return nil, err
	}
	cfg.Prometheus.URI = viper.GetString("prometheus.uri")
	cfg.Prometheus.Port = viper.GetInt("prometheus.port")

	return cfg, nil
}

func getEnvAsInt(name string) (int, error) {
	valueStr, err := getEnv(name)
	if valueStr == "" {
		return 0, err
	}
	value, err := strconv.Atoi(valueStr)
	if err == nil {
		return value, nil
	}
	return 0, fmt.Errorf("env var %v parsing failed: %v", name, err)
}

func getEnvAsString(name string) (string, error) {
	valueStr, err := getEnv(name)
	if valueStr == "" {
		return "", err
	}
	return valueStr, nil
}

func getEnvAsSlice(name string, sep string) ([]string, error) {
	valueStr, err := getEnv(name)
	if valueStr == "" {
		return nil, err
	}
	val := strings.Split(valueStr, sep)
	return val, nil
}

func getEnv(name string) (string, error) {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return "", fmt.Errorf("no such ENV variable: %v", name)
	}
	return valueStr, nil
}
