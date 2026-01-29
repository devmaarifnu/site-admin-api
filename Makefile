.PHONY: help run build test clean dev

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	go run cmd/main.go

build: ## Build the application
	go build -o site-admin-api cmd/main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build files
	rm -f site-admin-api
	rm -rf logs/

dev: ## Run with air (hot reload) - requires air to be installed
	air

install: ## Install dependencies
	go mod download
	go mod tidy

migrate-up: ## Run database migrations up
	@echo "Running migrations..."
	# Add migration tool command here

migrate-down: ## Run database migrations down
	@echo "Rolling back migrations..."
	# Add migration tool command here

lint: ## Run golangci-lint
	golangci-lint run

.DEFAULT_GOAL := help
