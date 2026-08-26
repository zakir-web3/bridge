#!/usr/bin/make -f

VERSION := $(shell git describe --tags --always --dirty 2>/dev/null)
ifeq ($(VERSION),)
VERSION := dev
endif
LDFLAGS := -s -w -X main.Version=$(VERSION)

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

e2e:
	bash scripts/e2e.sh
