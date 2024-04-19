SHELL:=/bin/bash

NAME:=goNDK
HASH:=$(shell git rev-parse --short HEAD)

all: clean lint test cover
.PHONY: all

clean:
	@echo "make clean"
	rm -f ./coverage.out
	rm -f ./lint.out

generate:
	@echo "make generate"
	go run github.com/mailru/easyjson/easyjson -all event/event.go
.PHONY: generate

lint:
	@echo "make lint"
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=5m ./... | tee lint.out
.PHONY: lint

test:
	@echo "make test"
	go test ./...
.PHONY: test

cover:
	@echo "make cover"
	go test -timeout=5m -coverprofile=coverage.out ./...
	./.meta/cover.sh
.PHONY: cover

docker.build:
	@echo "make docker.build"
	docker build . \
		-f Dockerfile -t $(NAME):$(HASH) \
		--build-arg BUILD_REVISION=$(HASH)
	docker tag $(NAME):$(HASH) $(NAME):latest
.PHONY: docker.build

docker.lint: docker.build
	@echo "make docker.lint"
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make lint
.PHONY: docker.lint

docker.test: docker.build
	@echo "make docker.test"
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make test
.PHONY: docker.test

docker.cover: docker.build
	@echo "make docker.cover"
	docker run --rm -v $(PWD):/app $(NAME):$(HASH) make cover
.PHONY: docker.cover

docker.shell: docker.build
	@echo "make docker.shell"
	docker run --rm -it -v $(PWD):/app $(NAME):$(HASH) /bin/bash
.PHONY: docker.shell
