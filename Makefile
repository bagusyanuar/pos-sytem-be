include .env
export

.PHONY: migrate-up migrate-down migrate-create seed run dev

# Migration commands
migrate-up:
	@echo "Running migrations..."
	@migrate -path database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	@echo "Rolling back migrations..."
	@migrate -path database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" down

migrate-force:
	@read -p "Enter version to force: " version; \
	migrate -path database/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" force $$version

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir database/migrations -seq $$name

# Seeder
seed:
	@echo "Running seeders..."
	@go run cmd/seeder/main.go

# Run application
run:
	@go run cmd/api/main.go

# Run with migration
run-migrate:
	@go run cmd/api/main.go -migrate

# Development with hot reload (using air)
dev:
	@air

# Install dependencies
install:
	@go mod download
	@go install github.com/cosmtrek/air@latest
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Database reset (down + up + seed)
db-reset: migrate-down migrate-up seed
	@echo "Database reset completed!"

# Setup project
setup:
	@echo "Setting up project..."
	@cp .env.example .env
	@mkdir -p storage/logs storage/temp
	@touch storage/logs/.gitkeep storage/temp/.gitkeep
	@go mod download
	@echo "Setup completed! Please update .env file with your configuration."

# Clean
clean:
	@echo "Cleaning..."
	@rm -rf tmp/
	@rm -rf storage/logs/*.log
	@rm -rf storage/temp/*
	@echo "Clean completed!"

# Test
test:
	@go test -v ./...

# Test with coverage
test-coverage:
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Build
build:
	@echo "Building application..."
	@go build -o bin/pos-api cmd/api/main.go
	@echo "Build completed! Binary: bin/pos-api"