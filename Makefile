.PHONY: help build install run test clean snapshot release dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the binary
	go build -o finfolio .

install: ## Install to $GOPATH/bin
	go install

run: ## Run the application
	go run main.go

dev: ## Run in development mode (example: make dev CMD="add")
	go run main.go $(CMD)

test: ## Run tests
	go test -v ./...

clean: ## Remove build artifacts
	rm -rf dist/
	rm -f finfolio

snapshot: ## Build snapshot release (no git tag required)
	goreleaser release --snapshot --clean

release: ## Create a new release (requires GITHUB_TOKEN)
	@echo "Current tags:"
	@git tag -l
	@echo ""
	@read -p "Enter new version (e.g., v0.2.0): " version; \
	git tag -a $$version -m "Release $$version" && \
	git push origin $$version && \
	goreleaser release --clean

tidy: ## Tidy go modules
	go mod tidy

fmt: ## Format code
	go fmt ./...

lint: ## Run linter (requires golangci-lint)
	golangci-lint run

check: fmt lint test ## Run all checks (format, lint, test)