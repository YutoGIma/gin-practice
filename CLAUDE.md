# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Development Commands

### Running the Application
```bash
# Start the application with database using Docker Compose
docker-compose up

# Build and start fresh
docker-compose up --build

# Run in detached mode
docker-compose up -d

# Stop all services
docker-compose down

# Clean stop with volume removal
docker-compose down -v
```

### Direct Go Commands
```bash
# Run the application directly (requires local PostgreSQL)
go run main.go

# Build the binary
go build -o gin-practice main.go

# Download dependencies
go mod download

# Clean up dependencies
go mod tidy
```

### Database Operations
```bash
# Run migrations via Docker
docker-compose run migrate

# Run migrations directly
go run app/migrate/migrate.go
```

### Viewing Logs
```bash
# Application logs
docker-compose logs -f app

# Database logs
docker-compose logs -f db
```

## Architecture Overview

This application follows a **Clean Architecture** pattern with unidirectional dependency flow:

```
HTTP Request → Router → Controller → UseCase → Service → Database
                ↑                                          ↓
                └────────── HTTP Response ←────────────────┘
```

### Layer Responsibilities

1. **Controller Layer** (`app/controller/`)
   - Handles HTTP requests/responses
   - Request validation and binding
   - Error response formatting
   - No business logic

2. **UseCase Layer** (`app/usecase/`)
   - Contains business logic and validation
   - Orchestrates service calls
   - Handles business rule enforcement
   - Transaction boundaries (if implemented)

3. **Service Layer** (`app/service/`)
   - Direct database operations via GORM
   - No business logic, only data access
   - Returns domain models

4. **Model Layer** (`app/model/`)
   - Domain entities with GORM tags
   - All models embed `BaseModel` for common fields (ID, timestamps, soft delete)
   - Relationships defined via GORM tags

### Key Architectural Patterns

**Aggregator Pattern**: The application uses `BaseController`, `BaseUseCase`, and `BaseService` structs to aggregate all controllers, use cases, and services respectively. This simplifies dependency injection in `main.go`.

**Dependency Injection**: Constructor-based injection flows from main.go:
```
DB → BaseService → BaseUseCase → BaseController → Router
```

**Error Handling**: Custom error types in `app/errors/` with predefined HTTP status codes:
- `NewValidationError()` → 400 Bad Request
- `NewNotFoundError()` → 404 Not Found
- `NewInternalError()` → 500 Internal Server Error

**Soft Deletes**: All models support soft deletion via GORM's `DeletedAt` field in `BaseModel`.

**Multi-tenancy**: The `Tenant` model represents different locations/branches. `Inventory` links products to specific tenants.

### Adding New Features

To add a new feature (e.g., "Order"):

1. Create model in `app/model/order_model.go`
2. Create service in `app/service/order_service.go`
3. Create use case in `app/usecase/order_usecase.go`
4. Create controller in `app/controller/order_controller.go`
5. Add to aggregators in respective `base_*.go` files
6. Wire up in `main.go`
7. Add routes in `app/routes/routes.go`

### Database Schema Management

- **Auto-migration** runs on startup via `db.AutoMigrate()` in main.go
- Models define schema through GORM tags
- Foreign keys and indexes are defined in model structs

### Environment Variables

```bash
DB_HOST=db          # Database host (default: "db" for Docker)
DB_PORT=5432        # Database port
DB_USER=user        # Database user
DB_PASSWORD=password # Database password
DB_NAME=myapp       # Database name
SERVER_PORT=8080    # Application port
ENV=development     # Environment (development/production)
```

### API Endpoints

- `GET/POST/PUT/DELETE /users` - User management
- `GET/POST/PUT/DELETE /products` - Product management
- `GET/POST/DELETE /inventories` - Inventory management
- `POST /inventories/purchase` - Purchase from inventory
- `GET/POST/PUT/DELETE /tenants` - Tenant management
- `GET /ping` - Health check

### Code Conventions from Existing Documentation

Based on `.cursor/rules/project.mdc`:

- **File naming**: Snake case (e.g., `user_controller.go`)
- **Package names**: Lowercase snake case
- **Struct names**: PascalCase (e.g., `UserController`)
- **Method names**: PascalCase (e.g., `GetUser`)
- **Variable names**: camelCase (e.g., `userID`)

**Import order**:
1. Internal packages (`myapp/app/...`)
2. External packages (`github.com/...`)
3. Standard library

### Testing

Currently, no test files exist in the project. When adding tests:
- Place test files next to the code they test
- Name test files with `_test.go` suffix
- Use table-driven tests for comprehensive coverage