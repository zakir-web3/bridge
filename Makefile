#!/usr/bin/make -f

COMMIT := $(shell git rev-parse --short=8 HEAD)
LDFLAGS := -s -w -X main.Version=$(COMMIT)

install:
	go install -mod=readonly -ldflags '$(LDFLAGS)'

build:
	go build -mod=readonly -ldflags '$(LDFLAGS)' -o ./bin/bridge

test:
	go test ./...

lint:
	golangci-lint run -v ./...

format:
	golangci-lint fmt -v ./...
