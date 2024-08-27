# Variables
APP_UNIQUE_NAME=go_rest_api
APP_NAME="Simple Go REST API"
FILE_NAME=main.go
CMD_PATH=./cmd
VERSION=1.0.0
CURRENT_DIR=$(shell pwd)
MIGRATIONS_PATH=/db/migrations
SEEDS_PATH=/db/seeds
#Colors
C_RED=\033[0;31m
END_COLOR=\e[0m
SQLITE_DB_PATH=/db/db.db
#ARGS
## Migration


# build: Build the application
build:
	@echo "Building: $(APP_UNIQUE_NAME)"
	@cd $(CMD_PATH) && go build -o  ../bin/${APP_UNIQUE_NAME} $(FILE_NAME)
	@echo "Build completed. Avaiable on $(CURRENT_DIR)/bin"

install:
	@echo "Installing: $(APP_UNIQUE_NAME) dependencies..."
	@go mod tidy
	@echo "Dependencies installed"
	@echo "Installing Database Migration Tool (Goose)..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Goose installed!"
	@echo "Installing Air..."
	@go install github.com/air-verse/air@latest
	@echo "Air installed!"
	@echo "Built with" "$(C_RED)<3$(END_COLOR)" "by @harry-fruit"


# migrate: Run the database migrations
db-migrate-up:
	@echo "Running database migrations..."
	@goose -dir .$(MIGRATIONS_PATH) sqlite3 $(CURRENT_DIR)/db/db.db up
	@echo "Database migrations completed"

db-migrate-down:
	@echo "Dropping database last migration..."
	@goose -dir .$(MIGRATIONS_PATH) sqlite3 $(CURRENT_DIR)/db/db.db down
	@echo "Migration dropped"

db-migrate-reset:
	@echo "Reset database migrations..."
	@goose -dir .$(MIGRATIONS_PATH) sqlite3 $(CURRENT_DIR)/db/db.db reset
	@echo "Database migrations reset"

db-create-migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Please provide a migration name"; \
		exit 1; \
	fi
	@echo "Creating migration: $(MIGRATION_NAME)"
	@goose -dir $(CURRENT_DIR)$(MIGRATIONS_PATH) create $(MIGRATION_NAME) sql

db-seed:
	@echo "Seeding db..."
	@goose -dir .$(SEEDS_PATH) -no-versioning sqlite3 $(CURRENT_DIR)/db/db.db up
	@echo "Seeding completed"

db-create-seed:
	@if [ -z "$(SEED_NAME)" ]; then \
		echo "Please provide a seed name"; \
		exit 1; \
	fi
	@echo "Creating seed: $(SEED_NAME)"
	@goose -dir $(CURRENT_DIR)$(SEEDS_PATH) create $(SEED_NAME) sql

swag-compile:
	@echo "Compiling swagger docs..."
	@swag init -g ./cmd/main.go --dir ./ --output ./api
	@echo "Swagger docs compiled"