# SnackNews Daily Service

This project implements a shared service for managing news and categories using the Gin framework in Go. It provides a RESTful API for interacting with news items and categories, along with middleware for request processing.

## Project Structure

- **cmd/server/main.go**: Entry point of the application. Initializes the Gin engine, sets up routes, and starts the server.
- **internal/api**: Contains API-related files.
  - **routes.go**: Defines the API routes and associates them with handler functions.
  - **middleware.go**: Contains middleware functions for request processing.
- **internal/handlers**: Contains handler functions for various endpoints.
  - **health.go**: Health check handler.
  - **news.go**: Handlers for managing news items.
  - **categories.go**: Handlers for managing news categories.
- **internal/service**: Contains business logic for news and categories.
  - **news_service.go**: Business logic for news items.
  - **category_service.go**: Business logic for categories.
- **internal/repository**: Handles database interactions.
  - **db.go**: Database connection and configuration.
  - **news_repo.go**: Functions for accessing and manipulating news data.
- **internal/models**: Defines data models used in the application.
- **internal/config**: Configuration loading and management.
- **pkg/errors**: Custom error types and utility functions for error handling.
- **migrations**: SQL statements for initializing the database schema.
- **configs**: Configuration settings in YAML format.
- **scripts**: Shell scripts for running database migrations.
- **Dockerfile**: Instructions for building a Docker image for the application.
- **Makefile**: Defines build and deployment commands.
- **.env.example**: Example of environment variables for configuration.
- **go.mod**: Module definition and dependencies.
- **go.sum**: Checksums for module dependencies.

## Getting Started

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd SnackNews-daily-backend
   ```

2. **Install dependencies**:
   ```
   go mod tidy
   ```

3. **Run the application**:
   ```
   go run cmd/server/main.go
   ```

4. **Access the API**:
   The API will be available at `http://localhost:8080`.

## API Endpoints

- **Health Check**: `GET /health`
- **News**: 
  - `GET /news`
  - `POST /news`
  - `PUT /news/:id`
  - `DELETE /news/:id`
- **Categories**:
  - `GET /categories`
  - `POST /categories`
