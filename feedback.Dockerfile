FROM golang:latest AS builder

COPY . /build
WORKDIR /build
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/ocp-feedback-api cmd/ocp-feedback-api/main.go


#app
FROM alpine:latest
COPY --from=builder /build/bin/ocp-feedback-api /ocp-feedback-api/
COPY --from=builder /build/configs/feedback_cfg.yml /ocp-feedback-api/
WORKDIR /ocp-feedback-api
CMD ["./ocp-feedback-api","-config_name=feedback_cfg","-config_path=."]
