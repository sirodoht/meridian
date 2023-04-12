.PHONY: all
all: lint serve

.PHONY: lint
lint:
	GOGC=off golangci-lint run

.PHONY: format
format:
	go fmt ./...

.PHONY: serve
serve:
	modd

.PHONY: build
build:
	go build -v -o meridian ./cmd/server/main.go

.PHONY: test
test:
	go test -v ./...
