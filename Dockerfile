# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o guard

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/guard /usr/local/bin/guard

ENTRYPOINT ["guard"]
