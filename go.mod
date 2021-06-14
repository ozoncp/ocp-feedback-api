module github.com/ozoncp/ocp-feedback-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v0.0.0-20200227202807-02e2044944cc // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	google.golang.org/grpc v1.38.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api => ./pkg/ocp-feedback-api

replace github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api => ./pkg/ocp-proposal-api
