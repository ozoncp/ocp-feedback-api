package config

import (
	"github.com/spf13/viper"
)

type General struct {
	Chunks int
}

type GRPC struct {
	Host string
	Port int
}

type Postgres struct {
	ConnString   string
	MaxOpenConns int
	MaxIdleConns int
}

type Prometheus struct {
	URI  string
	Host string
	Port int
}

type Kafka struct {
	Brokers []string
}

type Config struct {
	General    General
	GRPC       GRPC
	Postgres   Postgres
	Prometheus Prometheus
	Kafka      Kafka
}

func Read(name, path string) (*Config, error) {

	// TODO: set defaults
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	cfg := &Config{}

	cfg.General.Chunks = viper.GetInt("general.chunks")

	cfg.GRPC.Host = viper.GetString("grpc.host")
	cfg.GRPC.Port = viper.GetInt("grpc.port")

	cfg.Postgres.ConnString = viper.GetString("postgres.conn_string")
	cfg.Postgres.MaxOpenConns = viper.GetInt("postgres.max_open_conns")
	cfg.Postgres.MaxIdleConns = viper.GetInt("postgres.max_idle_conns")

	cfg.Kafka.Brokers = viper.GetStringSlice("kafka.brokers")

	cfg.Prometheus.URI = viper.GetString("prometheus.uri")
	cfg.Prometheus.Host = viper.GetString("prometheus.host")
	cfg.Prometheus.Port = viper.GetInt("prometheus.port")

	return cfg, nil
}
