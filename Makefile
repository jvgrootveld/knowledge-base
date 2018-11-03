# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=knowledge-base
BINARY_UNIX=$(BINARY_NAME)_unix

SOURCE="./"
TARGET="build/$(BINARY_NAME)"

.PHONY: build
build:
	$(GOBUILD) -o "${TARGET}" -v "${SOURCE}"

run: build
	./build/$(BINARY_NAME)

.PHONY: build-linux
# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v