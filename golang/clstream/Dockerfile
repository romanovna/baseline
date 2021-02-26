FROM golang:1.16

ENV GO111MODULE on

ADD . ./src

WORKDIR ./src
RUN go mod download &&CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags "-w -s" ./cmd/server/main.go

ENTRYPOINT ["/go/bin/main"]
EXPOSE 8080