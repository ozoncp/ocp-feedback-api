PHONY: build
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		mkdir -p pkg/ocp-feedback-api
		protoc -I vendor.protogen -I /usr/local/include  -I api/ocp-feedback-api\
				--go_out=pkg/ocp-feedback-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-feedback-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-feedback-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-feedback-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-feedback-api:swagger \
				api/ocp-feedback-api/feedback-service.proto \
				api/ocp-feedback-api/feedback-messages.proto 
		mv pkg/ocp-feedback-api/github.com/ozoncp/ocp-feedback-api/pkg/ocp-feedback-api/* pkg/ocp-feedback-api
		rm -rf pkg/ocp-feedback-api/github.com
		mkdir -p cmd/ocp-feedback-api

		mkdir -p pkg/ocp-proposal-api
		protoc -I vendor.protogen -I /usr/local/include  -I api/ocp-proposal-api\
				--go_out=pkg/ocp-proposal-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-proposal-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-proposal-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-proposal-api \
				--swagger_out=allow_merge=true,merge_file_name=ocp-proposal-api:swagger \
				api/ocp-proposal-api/proposal-service.proto \
				api/ocp-proposal-api/proposal-messages.proto
		mv pkg/ocp-proposal-api/github.com/ozoncp/ocp-feedback-api/pkg/ocp-proposal-api/* pkg/ocp-proposal-api
		rm -rf pkg/ocp-proposal-api/github.com
		mkdir -p cmd/ocp-proposal-api
				

PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-feedback-api cmd/ocp-feedback-api/main.go
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-proposal-api cmd/ocp-proposal-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/ocp-feedback-api/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-feedback-api
		cp api/ocp-feedback-api/* vendor.protogen/api/ocp-feedback-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u github.com/envoyproxy/protoc-gen-validate	
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate	

.PHONY: test
test:
		go test -race ./... -coverprofile=cover.out 

.PHONY: test-coverage
test-coverage:
		go tool cover -func cover.out | grep total | awk '{print $3}'

.PHONY: clean
clean:
		rm cover.out
		rm -rf bin		

.PHONY: feedback_server
feedback_server:
		go run cmd/ocp-feedback-api/main.go -config_name=feedback_cfg -config_path=configs

.PHONY: proposal_server
proposal_server:
		go run cmd/ocp-proposal-api/main.go -config_name=proposal_cfg -config_path=configs