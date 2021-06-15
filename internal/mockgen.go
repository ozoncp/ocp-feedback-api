package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-feedback-api/internal/repo BatchAdder
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-feedback-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/prommetrics_mock.go -package=mocks github.com/ozoncp/ocp-feedback-api/internal/prommetrics PromMetrics
