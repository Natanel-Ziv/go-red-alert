FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /oref-alerts-go ./cmd/alerts-proxy

FROM alpine:3.18
COPY --from=builder /oref-alerts-go /usr/local/bin/
EXPOSE 9001
ENTRYPOINT ["/usr/local/bin/oref-alerts-go"]
