APP_UNIQUE_NAME=go_rest_api
APP_NAME="Simple Go REST API"
FILE_NAME=main.go
CMD_PATH=./cmd
VERSION=1.0.0


# run: Run the application
run:
	@echo "Running: $(APP_UNIQUE_NAME)"
	@cd $(CMD_PATH) && go run $(FILE_NAME)

# build: Build the application
build:
	@echo "Building: $(APP_UNIQUE_NAME)"
	@cd $(CMD_PATH) && go build -o  ../bin/${APP_UNIQUE_NAME} $(FILE_NAME)
