FROM golang:latest AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/ocp-proposal-api cmd/ocp-proposal-api/main.go


#app
FROM alpine:latest
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait
COPY --from=builder /build/bin/ocp-proposal-api /ocp-proposal-api/
COPY --from=builder /build/configs/proposal.yml /ocp-proposal-api/
COPY --from=builder /build/swagger/ocp-proposal-api.swagger.json /ocp-proposal-api/swagger/
WORKDIR /ocp-proposal-api
CMD /wait ; ./ocp-proposal-api -config_name=proposal_cfg -config_path=.
