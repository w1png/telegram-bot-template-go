SHELL := /bin/zsh

.PHONY: run build lint

run:
	go run .

build:
	go build


lint:
	golangci-lint run

