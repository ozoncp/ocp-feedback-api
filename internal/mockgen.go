package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-feedback-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/alarmer_mock.go -package=mocks github.com/ozoncp/ocp-feedback-api/internal/alarmer Alarmer
