.DEFAULT_GOAL := build

build:
	go build -o ./git-release ./cmd/git-release
.PHONY: build

test:
	go test -v ./...
.PHONY: test

lint:
	golangci-lint run ./...
.PHONY: lint