.DEFAULT_GOAL := build

VERSION := $(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
GIT_HASH := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date -u | sed 's| |_|'g)

# git@gitlab.com:personal116/dotfiles.git

.PHONY: run
run: build
	./config

.PHONY: build
build:
	CGO_ENABLED=0 \
	go build \
		-v \
		-ldflags="-s -w -X main.GitHash=$(GIT_HASH) -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" \
		./cmd/config

.PHONY: install
install:
	CGO_ENABLED=0 \
	go build \
		-v \
		-ldflags="-s -w -X main.GitHash=$(GIT_HASH) -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" \
		./cmd/config

.PHONY: generate-release-artifacts
generate-release-artifacts:
	make build-release GOOS="darwin" GOARCH="arm64"
	make build-release GOOS="darwin" GOARCH="amd64"
	make build-release GOOS="linux" GOARCH="arm64"
	make build-release GOOS="linux" GOARCH="amd64"

.PHONY: build-release
build-release:
	CGO_ENABLED=0 \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	go build \
		-ldflags="-s -w -X main.GitHash=$(GIT_HASH) -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" \
		./cmd/config && mv config config_$(GOOS)-$(GOARCH)

.PHONY: gofmt
gofmt:
	gofmt -l -s -w ./cmd/config ./

.PHONY: mod
mod:
	go mod tidy

.PHONY: lint 
lint:
	docker run \
		--rm -it \
		-w /sources \
		-v `pwd`:/sources \
		golangci/golangci-lint:v1.53.3 \
		golangci-lint -c .golangci.yml run
