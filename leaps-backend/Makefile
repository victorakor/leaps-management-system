.PHONY: help build up down logs test clean deploy

# Variables
DOCKER_COMPOSE := docker-compose
GO := go
BINARY_NAME := leaps-api

help:
	@echo "LEAPS API - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make build          - Build Docker images"
	@echo "  make up             - Start all services"
	@echo "  make down           - Stop all services"
	@echo "  make logs           - View service logs"
	@echo "  make logs-api       - View API logs"
	@echo "  make logs-db        - View database logs"
	@echo ""
	@echo "Testing:"
	@echo "  make test           - Run all tests"
	@echo "  make test-unit      - Run unit tests"
	@echo "  make test-coverage  - Run tests with coverage"
	@echo ""
	@echo "Database:"
	@echo "  make db-shell       - Connect to database shell"
	@echo "  make db-backup      - Backup database"
	@echo "  make db-restore     - Restore database"
	@echo ""
	@echo "Deployment:"
	@echo "  make deploy         - Deploy application"
	@echo "  make clean          - Clean up containers and volumes"
	@echo ""

# Build Docker images
build:
	@echo "Building Docker images..."
	$(DOCKER_COMPOSE) build

# Start services
up:
	@echo "Starting services..."
	$(DOCKER_COMPOSE) up -d
	@echo "Services started!"
	@echo "API: http://localhost:8080"
	@echo "pgAdmin: http://localhost:5050"

# Stop services
down:
	@echo "Stopping services..."
	$(DOCKER_COMPOSE) down

# View logs
logs:
	$(DOCKER_COMPOSE) logs -f

logs-api:
	$(DOCKER_COMPOSE) logs -f api

logs-db:
	$(DOCKER_COMPOSE) logs -f postgres

# Run tests
test:
	@echo "Running tests..."
	$(GO) test -v ./...

test-unit:
	@echo "Running unit tests..."
	$(GO) test -v -short ./...

test-coverage:
	@echo "Running tests with coverage..."
	$(GO) test -v -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Database commands
db-shell:
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -d leaps_db

db-backup:
	@echo "Backing up database..."
	$(DOCKER_COMPOSE) exec -T postgres pg_dump -U postgres leaps_db > backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "Backup complete!"

db-restore:
	@echo "Restoring database..."
	@read -p "Enter backup file path: " backup_file; \
	$(DOCKER_COMPOSE) exec -T postgres psql -U postgres leaps_db < $$backup_file
	@echo "Restore complete!"

# Deploy
deploy:
	@echo "Deploying application..."
	bash scripts/deploy.sh

# Clean up
clean:
	@echo "Cleaning up..."
	$(DOCKER_COMPOSE) down -v
	rm -f coverage.out
	@echo "Cleanup complete!"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	cp .env.example .env
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE) up -d
	@echo "Development environment ready!"

# Format code
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run ./...

# Build binary
build-binary:
	@echo "Building binary..."
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o $(BINARY_NAME) .
	@echo "Binary built: $(BINARY_NAME)"

# Run locally (without Docker)
run:
	@echo "Running API locally..."
	$(GO) run main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod download
	$(GO) mod tidy

# Health check
health:
	@echo "Checking API health..."
	@curl -s http://localhost:8080/health | jq .

# Show status
status:
	@echo "Service Status:"
	@$(DOCKER_COMPOSE) ps
