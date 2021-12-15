# Make Environment
BINARY 	= main
BIN 	= bin
VERSION ?=?
COMMIT	= $(shell git rev-parse HEAD)
BRANCH	= $(shell git rev-parse --abbrev-ref HEAD)

# Symlink
GITLAB_USERNAME = moehiqbal
BUILD_DIR		= $(GOPATH)/src/gitlab.com/${GITLAB_USERNAME}/bin/${BINARY}
CURRENT_DIR		= $(shell pwd)
BUILD_DIR_LINK	= $(shell readlink ${BUILD_DIR})

# Test
test:
	go test -v -cover -covermade=atomic ./...

build:
	go build -o bin/${BINARY} main.go

docker-image:
	docker build -t golang-http-example .

docker-run:
	docker run -dit --name golang-http-example -p 3000 golang-http-example

link-prepare:
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

lint:
	./bin/golangci-lint run ./...