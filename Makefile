.DEFAULT_GOAL := bin

bin:
	go build -o ./git-release ./cmd/git-release
.PHONY: bin