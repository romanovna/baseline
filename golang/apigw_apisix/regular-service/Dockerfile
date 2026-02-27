FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /regular-service .

FROM alpine:3.20
COPY --from=builder /regular-service /regular-service
EXPOSE 8083
CMD ["/regular-service"]
