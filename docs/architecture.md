
# Hexagonal Architecture Documentation

## Overview

This project implements a Go backend using Hexagonal Architecture (Ports and Adapters pattern), which promotes separation of concerns and makes the system more testable and maintainable.

## Architecture Layers

### 1. Core Domain (`/internal/core/`)
- **Entities**: Business objects with identity (e.g., User)
- **Value Objects**: Immutable objects that describe characteristics (e.g., Email)
- **Domain Services**: Business logic that doesn't naturally fit in entities

### 2. Ports (`/internal/port/`)
- **Repository Interfaces**: Define contracts for data persistence
- **Service Interfaces**: Define contracts for external services

### 3. Adapters (`/internal/adapter/`)
- **Repository Implementations**: Concrete implementations of repository interfaces
- **Service Implementations**: Concrete implementations of service interfaces
- **REST Handlers**: HTTP API endpoints
- **Middleware**: Cross-cutting concerns (logging, auth, etc.)

### 4. Application Layer (`/internal/app/`)
- **Use Cases**: Application-specific business logic
- **Orchestration**: Coordinates between domain and adapters

### 5. Configuration (`/internal/config/`)
- **Environment Configuration**: Manages application settings
- **Dependency Injection**: Wires up dependencies

## Dependency Flow

```
External World -> Adapters -> Ports -> Application -> Domain
```

- Domain has no dependencies on outer layers
- Application depends only on Domain and Ports
- Adapters implement Ports and depend on external systems
- Configuration wires everything together

## Benefits

1. **Testability**: Easy to mock dependencies and test in isolation
2. **Flexibility**: Easy to swap implementations (e.g., database, external services)
3. **Maintainability**: Clear separation of concerns
4. **Scalability**: New features can be added without affecting existing code
