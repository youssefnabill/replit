
# Hexagonal Architecture Go Backend

A Go backend implementation using Hexagonal Architecture (Ports and Adapters pattern).

## Architecture

This project follows the hexagonal architecture pattern with clear separation of concerns:

- **Domain Layer**: Core business logic, entities, and value objects
- **Application Layer**: Use cases and application services
- **Ports**: Interfaces that define contracts
- **Adapters**: Implementations of ports (repositories, services, handlers)

## Project Structure

```
/cmd
  └── main.go                 # Application entry point

/internal
  ├── core/                   # Domain layer
  │   └── user/              # User domain
  │       ├── entity.go      # User entity
  │       ├── service.go     # Domain service
  │       └── valueobject.go # Value objects
  │
  ├── port/                  # Interfaces
  │   ├── repository/        # Repository interfaces
  │   └── service/           # Service interfaces
  │
  ├── adapter/               # Implementations
  │   ├── repository/        # Repository implementations
  │   ├── service/           # Service implementations
  │   ├── rest/              # HTTP handlers
  │   └── middleware/        # Middleware
  │
  ├── app/                   # Application layer
  │   └── user_app.go        # User application service
  │
  └── config/                # Configuration
      ├── config.go          # Configuration structure
      └── env.go             # Environment loading

/pkg                         # Shared utilities
/migrations                  # Database migrations
/docs                        # Documentation
/tests                       # Tests
```

## Getting Started

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```

3. Test the API:
   ```bash
   # Create a user
   curl -X POST http://localhost:5000/api/v1/users \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","name":"Test User"}'

   # List users
   curl http://localhost:5000/api/v1/users
   ```

## Running Tests

```bash
go test ./tests/...
```

## API Endpoints

- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users/:id` - Get user by ID
- `GET /api/v1/users` - List users with pagination

## Environment Variables

- `PORT` - Server port (default: 5000)
- `ENV` - Environment (development/production)
