# Go Clean Architecture Template

A production-ready Go reference architecture demonstrating domain decoupling and maintainable backend services. This architecture synthesizes enterprise best practices, utilizing [Database Access](https://go.dev/doc/tutorial/database-access) and [Web Service with Gin](https://go.dev/doc/tutorial/web-service-gin) as foundational models.

This project is designed to showcase structural engineering concepts, clean architecture patterns, and the creation of highly testable, concurrent backend systems.

## Technology Stack

| Category | Technology / Principle |
| --- | --- |
| **Language** | Go (Golang) 1.21+ |
| **Framework** | Gin (for HTTP routing) |
| **Database** | MariaDB with GORM (ORM) |
| **Architecture** | Clean Architecture, SOLID, Repository Pattern, Dependency Injection |
| **Testing** | Go `testing` package (Unit & Integration) |
| **CI/CD** | GitHub Actions |

## Architecture Overview

This project adapts the [evrone/go-clean-template](https://github.com/evrone/go-clean-template), but intentionally strips away heavy boilerplate and unnecessary external adapters. It is a simplified, lean version optimized for rapid feature delivery and our specific project scope, ensuring we maintain the strict domain decoupling of Clean Architecture without the bloat:

```text
├── cmd/app/           # Application entry point
├── config/            # Configuration management
├── internal/          # Private application code
│   ├── app/          # Application setup and dependency injection
│   ├── controller/   # HTTP handlers and routing
│   ├── entity/       # Business entities and DTOs
│   ├── repository/   # Data access layer
│   └── usecase/      # Business logic layer
└── pkg/              # Public packages that can be imported
    ├── app/          # Application utilities
    └── database/     # Database connection utilities
```

## Getting Started

### Prerequisites

* Go 1.21+
* MariaDB (or Docker provisioning)

### Installation & Running

1. **Clone the repository:**
   ```bash
   git clone https://github.com/komangsidhiartha/go-learning-clean-arch.git
   cd go-learning-clean-arch
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   go run cmd/app/main.go
   ```

4. **(Optional) Run with database migrations:**
   ```bash
   go run -tags migrate cmd/app/main.go
   ```

The server will initialize on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/v1/albums` | Retrieves a list of all albums. |
| `GET` | `/v1/albums/{id}` | Retrieves a single album by ID. |
| `POST` | `/v1/albums` | Creates a new album. |

## Key Architectural Concepts

### 1. Clean Architecture Layers

* **Entities** (`internal/entity/`): Pure business models with zero external dependencies.
* **Use Cases** (`internal/usecase/`): Core business logic orchestrating repositories, strictly isolated from HTTP or database awareness.
* **Repositories** (`internal/repository/`): Data access abstractions implementing interfaces defined by the business layer.
* **Controllers** (`internal/controller/`): HTTP request handling, input sanitization, and delegation to use cases.

### 2. Dependency Injection

The application utilizes constructor-based dependency injection within `internal/app/app.go` to strictly decouple components. Repositories are injected into use cases, and use cases into controllers, guaranteeing that every architectural layer remains independently testable.

### 3. Interface-Based Design

Business logic relies entirely on interfaces rather than concrete implementations. This enforces modularity, allowing for frictionless mocking during test cycles and seamless swapping of underlying infrastructure (e.g., database migrations).

```go
// internal/usecase/type.go
type AlbumRepository interface {
    GetAlbums() ([]entity.Album, error)
    // ... other methods
}

// internal/usecase/function.go
func (uc *AlbumUseCase) GetAlbums() ([]entity.Album, error) {
    return uc.repo.GetAlbums() // Executes via any AlbumRepository implementation
}
```

## Testing Strategy

The isolated architecture guarantees frictionless testing pipelines:

* **Unit Tests:** Use cases validate business logic using mocked repositories.
* **Integration Tests:** Repositories execute against a live test database to ensure data persistence integrity.
* **HTTP Tests:** Controllers validate routing and input parsing using mocked use cases.

Execute the full test suite:

```bash
go test ./... -v
```
