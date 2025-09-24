# URL Shortener Service

A simple URL shortener service with basic analytics, built with Go, Gin, and Postgres.

## Features

* Shorten long URLs
* Redirect using short URLs
* Track redirect counts
* Dockerized for easy deployment
* GitHub Actions CI for linting and tests
* Optional Kubernetes deployment

## Prerequisites

* Go 1.23+
* Docker & Docker Compose
* (Optional) Kubernetes cluster and `kubectl`
* Postgres (if not using Docker Compose)

## Running Locally

```bash
git clone https://github.com/thangnguyen19801/url-shortener.git
cd url-shortener
export DATABASE_URL="postgres://shortener:shortener@localhost:5432/shortener?sslmode=disable"
go run cmd/server/main.go
```

Server will run at: `http://localhost:8080`

## Running with Docker

```bash
docker build -t url-shortener .
docker run -p 8080:8080 -e DATABASE_URL="postgres://shortener:shortener@host.docker.internal:5432/shortener?sslmode=disable" url-shortener
```

## Running with Docker Compose

```bash
docker-compose up --build
```

## GitHub Actions CI

Workflow `.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: 1.23

      - name: Download Go modules
        run: go mod download

      - name: Install golangci-lint
        run: |
          go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.0
          echo "$HOME/go/bin" >> $GITHUB_PATH

      - name: Run golangci-lint
        run: golangci-lint run ./...

      - name: Run tests
        run: go test ./... -v

```

## Kubernetes Deployment (Optional)

Apply deployment and service:

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Access via port-forward or LoadBalancer IP.

## License

MIT License
