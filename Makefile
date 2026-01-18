include .env
export

# Variables
APP_NAME=pos-api
BINARY_DIR=bin
MIGRATION_DIR=database/migrations
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Colors for output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m
COLOR_BLUE=\033[34m

.PHONY: help migrate-up migrate-down migrate-force migrate-create migrate-version seed run run-migrate dev install db-reset db-fresh setup clean test test-coverage build build-linux build-windows build-mac lint fmt check

# Default target
.DEFAULT_GOAL := help

## help: Show this help message
help:
	@echo "$(COLOR_BOLD)Available commands:$(COLOR_RESET)"
	@echo ""
	@echo "$(COLOR_GREEN)Development:$(COLOR_RESET)"
	@echo "  make dev              - Run with hot reload (Air)"
	@echo "  make run              - Run application"
	@echo "  make run-migrate      - Run application with migration"
	@echo ""
	@echo "$(COLOR_GREEN)Database:$(COLOR_RESET)"
	@echo "  make migrate-up       - Run migrations"
	@echo "  make migrate-down     - Rollback migrations"
	@echo "  make migrate-force    - Force migration version"
	@echo "  make migrate-create   - Create new migration"
	@echo "  make migrate-version  - Show current migration version"
	@echo "  make seed             - Run seeders"
	@echo "  make db-reset         - Reset database (down + up + seed)"
	@echo "  make db-fresh         - Fresh database (drop + up + seed)"
	@echo ""
	@echo "$(COLOR_GREEN)Build:$(COLOR_RESET)"
	@echo "  make build            - Build application"
	@echo "  make build-linux      - Build for Linux"
	@echo "  make build-windows    - Build for Windows"
	@echo "  make build-mac        - Build for macOS"
	@echo ""
	@echo "$(COLOR_GREEN)Testing:$(COLOR_RESET)"
	@echo "  make test             - Run tests"
	@echo "  make test-coverage    - Run tests with coverage"
	@echo ""
	@echo "$(COLOR_GREEN)Code Quality:$(COLOR_RESET)"
	@echo "  make lint             - Run linter"
	@echo "  make fmt              - Format code"
	@echo "  make check            - Run fmt + lint + test"
	@echo ""
	@echo "$(COLOR_GREEN)Setup & Clean:$(COLOR_RESET)"
	@echo "  make install          - Install dependencies"
	@echo "  make setup            - Setup project"
	@echo "  make clean            - Clean temporary files"

## migrate-up: Run all pending migrations
migrate-up:
	@echo "$(COLOR_BLUE)Running migrations...$(COLOR_RESET)"
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up
	@echo "$(COLOR_GREEN)✓ Migrations completed!$(COLOR_RESET)"

## migrate-down: Rollback last migration
migrate-down:
	@echo "$(COLOR_YELLOW)Rolling back migrations...$(COLOR_RESET)"
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1
	@echo "$(COLOR_GREEN)✓ Rollback completed!$(COLOR_RESET)"

## migrate-down-all: Rollback all migrations
migrate-down-all:
	@echo "$(COLOR_YELLOW)Rolling back all migrations...$(COLOR_RESET)"
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down
	@echo "$(COLOR_GREEN)✓ All migrations rolled back!$(COLOR_RESET)"

## migrate-force: Force migration to specific version
migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" force $$version

## migrate-create: Create new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$name && \
	echo "$(COLOR_GREEN)✓ Migration created!$(COLOR_RESET)"

## migrate-version: Show current migration version
migrate-version:
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" version

## seed: Run database seeders
seed:
	@echo "$(COLOR_BLUE)Running seeders...$(COLOR_RESET)"
	@go run cmd/seeder/main.go
	@echo "$(COLOR_GREEN)✓ Seeding completed!$(COLOR_RESET)"

## run: Run application
run:
	@echo "$(COLOR_BLUE)Starting application...$(COLOR_RESET)"
	@go run cmd/api/main.go

## run-migrate: Run application with migration
run-migrate:
	@echo "$(COLOR_BLUE)Starting application with migration...$(COLOR_RESET)"
	@go run cmd/api/main.go -migrate

## dev: Run with hot reload using Air
dev:
	@echo "$(COLOR_BLUE)Starting development server with hot reload...$(COLOR_RESET)"
	@air

## install: Install all dependencies and tools
install:
	@echo "$(COLOR_BLUE)Installing dependencies...$(COLOR_RESET)"
	@go mod download
	@go install github.com/air-verse/air@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(COLOR_GREEN)✓ Dependencies installed!$(COLOR_RESET)"

## db-reset: Reset database (rollback all, migrate up, seed)
db-reset: migrate-down-all migrate-up seed
	@echo "$(COLOR_GREEN)✓ Database reset completed!$(COLOR_RESET)"

## db-fresh: Fresh database (force version 0, migrate up, seed)
db-fresh:
	@echo "$(COLOR_YELLOW)Resetting database...$(COLOR_RESET)"
	@migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" drop -f
	@$(MAKE) migrate-up
	@$(MAKE) seed
	@echo "$(COLOR_GREEN)✓ Fresh database completed!$(COLOR_RESET)"

## setup: Initial project setup
setup:
	@echo "$(COLOR_BLUE)Setting up project...$(COLOR_RESET)"
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "$(COLOR_GREEN)✓ .env file created$(COLOR_RESET)"; \
	else \
		echo "$(COLOR_YELLOW)⚠ .env already exists$(COLOR_RESET)"; \
	fi
	@mkdir -p storage/logs storage/temp tmp
	@touch storage/logs/.gitkeep storage/temp/.gitkeep
	@go mod download
	@echo "$(COLOR_GREEN)✓ Setup completed! Please update .env file with your configuration.$(COLOR_RESET)"

## clean: Clean temporary files and build artifacts
clean:
	@echo "$(COLOR_BLUE)Cleaning...$(COLOR_RESET)"
	@rm -rf tmp/
	@rm -rf $(BINARY_DIR)/
	@rm -rf storage/logs/*.log
	@rm -rf storage/temp/*
	@rm -f build-errors.log
	@rm -f coverage.out coverage.html
	@echo "$(COLOR_GREEN)✓ Clean completed!$(COLOR_RESET)"

## test: Run all tests
test:
	@echo "$(COLOR_BLUE)Running tests...$(COLOR_RESET)"
	@go test -v -race ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "$(COLOR_BLUE)Running tests with coverage...$(COLOR_RESET)"
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(COLOR_GREEN)✓ Coverage report generated: coverage.html$(COLOR_RESET)"
	@go tool cover -func=coverage.out

## build: Build application for current OS
build:
	@echo "$(COLOR_BLUE)Building application...$(COLOR_RESET)"
	@mkdir -p $(BINARY_DIR)
	@go build -ldflags="-s -w" -o $(BINARY_DIR)/$(APP_NAME) cmd/api/main.go
	@echo "$(COLOR_GREEN)✓ Build completed! Binary: $(BINARY_DIR)/$(APP_NAME)$(COLOR_RESET)"

## build-linux: Build for Linux
build-linux:
	@echo "$(COLOR_BLUE)Building for Linux...$(COLOR_RESET)"
	@mkdir -p $(BINARY_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY_DIR)/$(APP_NAME)-linux-amd64 cmd/api/main.go
	@echo "$(COLOR_GREEN)✓ Linux build completed!$(COLOR_RESET)"

## build-windows: Build for Windows
build-windows:
	@echo "$(COLOR_BLUE)Building for Windows...$(COLOR_RESET)"
	@mkdir -p $(BINARY_DIR)
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY_DIR)/$(APP_NAME)-windows-amd64.exe cmd/api/main.go
	@echo "$(COLOR_GREEN)✓ Windows build completed!$(COLOR_RESET)"

## build-mac: Build for macOS
build-mac:
	@echo "$(COLOR_BLUE)Building for macOS...$(COLOR_RESET)"
	@mkdir -p $(BINARY_DIR)
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY_DIR)/$(APP_NAME)-darwin-amd64 cmd/api/main.go
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BINARY_DIR)/$(APP_NAME)-darwin-arm64 cmd/api/main.go
	@echo "$(COLOR_GREEN)✓ macOS builds completed!$(COLOR_RESET)"

## build-all: Build for all platforms
build-all: build-linux build-windows build-mac
	@echo "$(COLOR_GREEN)✓ All builds completed!$(COLOR_RESET)"

## lint: Run linter
lint:
	@echo "$(COLOR_BLUE)Running linter...$(COLOR_RESET)"
	@golangci-lint run ./...

## fmt: Format code
fmt:
	@echo "$(COLOR_BLUE)Formatting code...$(COLOR_RESET)"
	@go fmt ./...
	@gofmt -s -w .
	@echo "$(COLOR_GREEN)✓ Code formatted!$(COLOR_RESET)"

## check: Run format, lint, and tests
check: fmt lint test
	@echo "$(COLOR_GREEN)✓ All checks passed!$(COLOR_RESET)"

## docker-build: Build Docker image
docker-build:
	@echo "$(COLOR_BLUE)Building Docker image...$(COLOR_RESET)"
	@docker build -t $(APP_NAME):latest .
	@echo "$(COLOR_GREEN)✓ Docker image built!$(COLOR_RESET)"

## docker-run: Run Docker container
docker-run:
	@echo "$(COLOR_BLUE)Running Docker container...$(COLOR_RESET)"
	@docker run -p 8080:8080 --env-file .env $(APP_NAME):latest

## docker-compose-up: Start all services with docker-compose
docker-compose-up:
	@docker-compose up -d

## docker-compose-down: Stop all services
docker-compose-down:
	@docker-compose down

## mod-tidy: Tidy go modules
mod-tidy:
	@echo "$(COLOR_BLUE)Tidying modules...$(COLOR_RESET)"
	@go mod tidy
	@echo "$(COLOR_GREEN)✓ Modules tidied!$(COLOR_RESET)"

## mod-vendor: Vendor dependencies
mod-vendor:
	@echo "$(COLOR_BLUE)Vendoring dependencies...$(COLOR_RESET)"
	@go mod vendor
	@echo "$(COLOR_GREEN)✓ Dependencies vendored!$(COLOR_RESET)"