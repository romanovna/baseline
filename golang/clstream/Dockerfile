ARG alpine_version=3.13
ARG go_version=1.16
FROM golang:$go_version-alpine$alpine_version AS builder

ENV GO111MODULE on

ADD . ./src

WORKDIR ./src
RUN go mod download &&CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags "-w -s" ./cmd/server/main.go

FROM alpine:$alpine_version
WORKDIR /app

COPY --from=builder /go/bin/main /app/
COPY --from=builder /go/src/server_config.yaml /app

ENTRYPOINT ["/app/main"]
EXPOSE 8080