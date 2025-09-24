# Build stage
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o url-service ./cmd/server

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /app/url-shortener .

EXPOSE 8080

CMD ["./url-shortener"]
