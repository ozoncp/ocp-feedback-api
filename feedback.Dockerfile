FROM golang:latest AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/ocp-feedback-api cmd/ocp-feedback-api/main.go


#app
FROM alpine:latest
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
COPY --from=builder /build/bin/ocp-feedback-api /ocp-feedback-api/
COPY --from=builder /build/configs/feedback_cfg.yml /ocp-feedback-api/
COPY --from=builder /build/swagger/ocp-feedback-api.swagger.json /ocp-feedback-api/swagger/
WORKDIR /ocp-feedback-api
CMD /wait ; ./ocp-feedback-api -config_name=feedback_cfg -config_path=.

