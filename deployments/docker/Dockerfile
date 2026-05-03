FROM golang:alpine AS base
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Live reloading with air
FROM base AS dev

RUN go install github.com/air-verse/air@v1.65.1
CMD ["air", "-c", ".air.toml"]

# Compile binary
FROM base AS builder

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/chart-gator-api ./cmd/api

# Production
FROM alpine:3.23.4 AS prod
WORKDIR /root/

COPY --from=builder /app/chart-gator-api .
COPY --from=builder /app/favicon.ico .

EXPOSE 8080
CMD ["./chart-gator-api"]
