FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /oref-alerts-go ./cmd/alerts-proxy

FROM alpine:3.18
COPY --from=builder /oref-alerts-go /usr/local/bin/
EXPOSE 9001
ENTRYPOINT ["/usr/local/bin/oref-alerts-go"]
