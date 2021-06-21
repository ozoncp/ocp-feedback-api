FROM golang:latest AS builder

COPY . /build
WORKDIR /build
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/ocp-proposal-api cmd/ocp-proposal-api/main.go


#app
FROM alpine:latest
COPY --from=builder /build/bin/ocp-proposal-api /ocp-proposal-api/
COPY --from=builder /build/configs/proposal.yml /ocp-proposal-api/
WORKDIR /ocp-proposal-api
CMD ["./ocp-proposal-api","-config_name=proposal_cfg","-config_path=."]
