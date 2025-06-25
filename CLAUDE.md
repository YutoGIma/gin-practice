# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Quick Start

```bash
# First time setup
make reset        # Clean database and start fresh

# Daily development
make dev          # Start with logs
make test-v       # Run tests with details
make swagger      # Generate API docs
```

## Common Development Commands

### Using Makefile (Recommended)

```bash
# Show all available commands
make help

# Container Management
make build        # Build Docker images
make up           # Start containers in background
make down         # Stop and remove containers
make start        # Start existing containers
make stop         # Stop containers (no removal)
make restart      # Restart containers
make status       # Show detailed status

# Development
make dev          # Development mode (build + logs)
make run          # Run in foreground
make logs         # View all logs
make logs-app     # View application logs only

# Database
make reset        # Full reset (clean + up + migrate)
make migrate      # Run database migrations
make seed         # Load seed data (currently automatic)
make db-shell     # Connect to PostgreSQL

# Testing & Quality
make test         # Run tests
make test-v       # Run tests with verbose output
make lint         # Run golangci-lint
make fmt          # Format code
make vet          # Run go vet
make tidy         # Clean up go.mod

# API Documentation
make swagger      # Generate Swagger docs
make swagger-fmt  # Format Swagger comments

# Cleanup
make clean        # Remove containers and volumes
make clean-all    # Remove all Docker resources
make prune        # Remove unused Docker resources

# Debug
make shell        # Enter app container shell
make status       # Show container and volume status
```

### Key Makefile Commands

| Command | Description | When to Use |
|---------|-------------|-------------|
| `make reset` | Full database reset | Starting fresh or fixing DB issues |
| `make dev` | Development mode | Active development with live logs |
| `make test-v` | Verbose tests | Debugging test failures |
| `make swagger` | Generate API docs | After adding/updating endpoints |
| `make clean` | Cleanup | Before switching branches or projects |

## Architecture Overview

This application follows a **Clean Architecture** pattern with unidirectional dependency flow:

```
HTTP Request → Router → Controller → UseCase → Service → Model → Database
                ↑                                                      ↓
                └────────────── HTTP Response ←────────────────────────┘
```

### Layer Responsibilities

1. **Router Layer** (`app/routes/`)
   - URL routing and middleware setup
   - Endpoint grouping and versioning
   - Swagger endpoint registration

2. **Controller Layer** (`app/controller/`)
   - HTTP request/response handling
   - Request validation and binding
   - Error response formatting
   - Swagger documentation

3. **UseCase Layer** (`app/usecase/`)
   - Business logic implementation
   - Transaction management
   - Input validation via validators
   - Error handling

4. **Service Layer** (`app/service/`)
   - Data access logic
   - Database queries
   - External service integration

5. **Model Layer** (`app/model/`)
   - Domain entities
   - Database schema definitions
   - Business rules

6. **Validation Layer** (`app/usecase/validation/`)
   - Centralized validation logic
   - Reusable validation methods
   - Domain-specific validators

## Key Patterns and Conventions

### Dependency Injection
- Constructor-based injection
- Aggregated through Base structs (BaseController, BaseUseCase, BaseService)
- Clean separation of concerns

### Error Handling
- Custom error types in `app/errors/`
- Middleware-based error handling
- Consistent error response format

### Database
- GORM with PostgreSQL
- Soft deletes enabled
- Auto-migration with seed data
- Upsert pattern for seed data (prevents duplicates)

### API Documentation
- Swagger/OpenAPI integration
- Inline documentation in controllers
- Auto-generated via swag tool
- Available at `/swagger/index.html`

### Testing Approach
```bash
# Run all tests
make test

# Run specific package tests
docker-compose run --rm app go test ./app/usecase/... -v

# Run with coverage
docker-compose run --rm app go test -cover ./...
```

## Project Structure

```
gin-practice/
├── app/
│   ├── controller/     # HTTP handlers
│   ├── usecase/        # Business logic
│   │   ├── request/    # Request DTOs
│   │   └── validation/ # Validators
│   ├── service/        # Data access
│   ├── model/          # Domain models
│   ├── routes/         # Router setup
│   ├── middleware/     # HTTP middleware
│   ├── errors/         # Custom errors
│   └── infra/          # Infrastructure
│       └── seed/       # Database seeds
├── docs/               # Swagger files
├── Makefile           # Development commands
├── docker-compose.yml  # Container setup
├── Dockerfile         # App container
└── main.go            # Entry point
```

## Recent Features

### Price Setting System
- Store-specific pricing with sale functionality
- Time-based pricing with start/end dates
- Automatic price activation/deactivation
- API endpoints under `/inventories/:id/prices`

### Order Management
- Complete order workflow with status tracking
- Inventory management integration
- Transaction support for data consistency
- User order history at `/users/:id/orders`

### Validation Framework
- Centralized validation in `app/usecase/validation/`
- Base validator with common methods
- Domain-specific validators for each entity
- Consistent error messages in Japanese

## Environment Variables

```bash
# Database Configuration (in docker-compose.yml)
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=myapp

# Application Configuration
GIN_MODE=debug  # Set to "release" for production
```

## Common Tasks

### Adding a New Feature
1. Create model in `app/model/`
2. Add to `GetModels()` in `base_model.go`
3. Create service in `app/service/`
4. Create usecase in `app/usecase/`
5. Create controller in `app/controller/`
6. Add routes in `app/routes/routes.go`
7. Update Base aggregators
8. Generate Swagger docs: `make swagger`

### Debugging Database Issues
```bash
# Check current data
make db-shell
\dt              # List tables
\d table_name    # Describe table
SELECT * FROM table_name;

# Reset if needed
make reset       # Full reset
```

### Troubleshooting

**Container won't start:**
```bash
make clean-all   # Remove everything
make reset       # Fresh start
```

**Port already in use:**
```bash
# Check what's using port 8080
lsof -i :8080
# Kill the process or change port in docker-compose.yml
```

**Database connection errors:**
```bash
make down
make clean
make up
make logs  # Check for specific errors
```

## Best Practices

1. **Always use Makefile commands** instead of raw docker-compose
2. **Run tests before committing**: `make test`
3. **Keep Swagger docs updated**: `make swagger` after API changes
4. **Use proper error handling** with custom error types
5. **Follow existing patterns** for consistency
6. **Use transactions** for operations affecting multiple tables
7. **Implement proper validation** in the validation layer