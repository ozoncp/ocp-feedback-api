module github.com/ozoncp/ocp-feedback-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/Shopify/sarama v1.29.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.3.0 // indirect
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc // indirect
	github.com/spf13/viper v1.8.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
	golang.org/x/net v0.0.0-20210610132358-84b48f89b13b // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	google.golang.org/grpc v1.38.0
)

replace github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api => ./pkg/ocp-feedback-api

replace github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api => ./pkg/ocp-proposal-api
