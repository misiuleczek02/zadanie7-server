.PHONY: build run test lint fmt hooks

build:
	go build -o bin/server ./

run:
	go run ./

test:
	go test -coverprofile=coverage.out ./...

lint:
	golangci-lint run ./...

fmt:
	gofmt -w .
	goimports -w .

hooks:
	bash scripts/install-hooks.sh
