.PHONY: build run clean test lint help

BINARY_NAME=mp3-organizer

build: ## Build the binary
	go build -o $(BINARY_NAME) ./cmd/mp3-tag-to-folder

run: ## Build and run the binary
	go run ./cmd/mp3-tag-to-folder/main.go

clean: ## Remove build artifacts
	go clean
	rm -f $(BINARY_NAME)
	rm -rf output/

test: ## Run tests
	go test -v ./...

lint: ## Run golangci-lint
	golangci-lint run

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
