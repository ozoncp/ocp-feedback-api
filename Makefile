PHONY: build
build: .generate

PHONY: .generate
.generate:
				mkdir -p pkg/ocp-feedback-api
				protoc \
					--go_out=pkg/ocp-feedback-api --go_opt=paths=import \
 				    --go-grpc_out=pkg/ocp-feedback-api --go-grpc_opt=paths=import \
					api/ocp-feedback-api/ocp-feedback-api.proto
				mv pkg/ocp-feedback-api/github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api/* pkg/ocp-feedback-api
				rm -rf pkg/ocp-feedback-api/github.com
