FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /client-service .

FROM alpine:3.20
COPY --from=builder /client-service /client-service
EXPOSE 8084
CMD ["/client-service"]
