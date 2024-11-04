.PHONY: build
build:
	go build -v ./core/api/

.PHONY: test
test:
	SET CGO_ENABLED=1&& go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build