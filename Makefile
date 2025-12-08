# Makefile

# Define the Go binary name
BINARY_NAME=news-shared-service

# Define the Go build command
build:
	go build -o $(BINARY_NAME) ./cmd/server

# Define the command to run the application
run: build
	./$(BINARY_NAME)

# Define the command to clean up the build artifacts
clean:
	rm -f $(BINARY_NAME)

# Define the command to run database migrations
migrate:
	sh scripts/migrate.sh

# Define the command to run tests
test:
	go test ./...

# Define the command to format the code
fmt:
	go fmt ./...

# Define the command to install dependencies
deps:
	go mod tidy

# Define the default target
.PHONY: build run clean migrate test fmt deps
.DEFAULT_GOAL := run