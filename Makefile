# Variables
APP_UNIQUE_NAME=go_rest_api
APP_NAME="Simple Go REST API"
FILE_NAME=main.go
CMD_PATH=./cmd
VERSION=1.0.0
CURRENT_DIR=$(shell pwd)
MIGRATIONS_PATH=/db/migrations
#Colors
C_RED=\033[0;31m
END_COLOR=\e[0m
SQLITE_DB_PATH=/db/db.db
#ARGS
## Migration


# run: Run the application
run:
	@echo "Running: $(APP_UNIQUE_NAME)"
	@cd $(CMD_PATH) && go run $(FILE_NAME)

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
	@echo "Goose installed - You are ready to go! Execute 'make run' to start the application :D"
	@echo "Built with" "$(C_RED)<3$(END_COLOR)" "by @harry-fruit"


# migrate: Run the database migrations
db-migrate:
	@echo "Running database migrations..."
	@goose -dir $(MIGRATIONS_PATH) sqlite3 $(CURRENT_DIR)/db/db.db up
	@echo "Database migrations completed"

db-create-migration:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Please provide a migration name"; \
		exit 1; \
	fi
	@echo "Creating migration: $(MIGRATION_NAME)"
	@goose -dir $(CURRENT_DIR)$(MIGRATIONS_PATH) create $(MIGRATION_NAME) sql