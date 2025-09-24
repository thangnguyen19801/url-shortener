.PHONY: build docker run test fmt lint

build:
	go build -o bin/url-shortener ./cmd/server

docker:
	docker build -t yourdockerhubusername/url-shortener:latest .

run:
	go run ./cmd/server

test:
	go test ./... -v

fmt:
	gofmt -w .

lint:
	golangci-lint run
