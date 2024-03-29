SHELL:=/bin/bash

NAME:=gondo
HASH:=$(shell git rev-parse --short HEAD)

all: test cover
.PHONY: all

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=5m ./... | tee lint.out
.PHONY: lint

test:
	go test ./...
.PHONY: test

cover:
	@go test \
		-timeout=5m \
		-coverprofile=coverage.out \
		-covermode=atomic \
		-coverpkg $(go list github.com/niallyoung/gondo/...) \
		./...
	$(./.meta/cover.sh)
.PHONY: cover

docker.build:
	docker build . \
		-f Dockerfile -t $(NAME):$(HASH) \
		--build-arg BUILD_REVISION=$(HASH)
	docker tag $(NAME):$(HASH) $(NAME):latest
.PHONY: docker.build

docker.lint: docker.build
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make lint
.PHONY: docker.lint

docker.test: docker.build
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make test
.PHONY: docker.test

docker.cover: docker.build
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make cover
.PHONY: docker.cover

docker.shell: docker.build
	docker run --rm -it -v $(PWD):/app $(NAME):$(HASH) /bin/bash
.PHONY: docker.shell
