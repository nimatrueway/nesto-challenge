SHELL = /usr/bin/env bash

deps-macos:
	# install golang
	brew install go
	# install colima for running docker in macOS
	brew install colima
	colima start --mount-type 9p
	# used for generating mocks for interfaces
	go install github.com/vektra/mockery/v2@v2.43.2
	# used for linting and static analysis
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

mock-database:
	pushd ../db-migrations && \
	docker build -t db-migrations . && docker run -p 5432:5432 db-migrations; \
	popd

lint: deps-macos
	golangci-lint run

test:
	mockery
	go test ./...

run-db-migrations:
	go run ./cmd/migration

run:
	go run ./cmd/server

.PHONY: deps lint test run-db-migrations run