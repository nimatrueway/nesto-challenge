SHELL = /usr/bin/env bash

deps:
	# used for generating mocks for interfaces
	go install github.com/vektra/mockery/v2@v2.43.2
	# used for linting and static analysis
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

lint: deps
	golangci-lint run

test:
	mockery
	go test ./...

run-db-migrations:
	go run ./cmd/migration

run:
	go run ./cmd/server

.PHONY: deps lint test run-db-migrations run