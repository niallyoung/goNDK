.PHONY: build test clean install

build:
	go build -o giftwrap .

test:
	go test ./...

test-short:
	go test ./... -short

clean:
	rm -f giftwrap

install: build
	cp giftwrap $(HOME)/bin/giftwrap

run-send:
	./giftwrap send --help

run-receive:
	./giftwrap receive --help
