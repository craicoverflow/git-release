.DEFAULT_GOAL := bin

bin:
	go build -o ./git-releaser ./cmd/git-releaser
.PHONY: bin