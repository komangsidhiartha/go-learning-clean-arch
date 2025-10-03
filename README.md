# Go Clean Architecture Template

A learning-focused Go template demonstrating clean architecture principles by combining two official Go tutorials: [Database Access](https://go.dev/doc/tutorial/database-access) and [Web Service with Gin](https://go.dev/doc/tutorial/web-service-gin).

This project is designed to showcase best practices in Go project templating, architecture patterns, and creating testable, maintainable backend services.

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

This project follows the [evrone/go-clean-template](https://github.com/evrone/go-clean-template) approach with a simplified structure for clarity:

```
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
* MariaDB (or Docker to run it)

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

The server will start on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| `GET` | `/v1/albums` | Retrieves a list of all albums. |
| `GET` | `/v1/albums/{id}` | Retrieves a single album by ID. |
| `POST` | `/v1/albums` | Creates a new album. |

## Key Architectural Concepts

### 1. Clean Architecture Layers

* **Entities** (`internal/entity/`): Pure business models with no external dependencies.
* **Use Cases** (`internal/usecase/`): Business logic that orchestrates repositories, with no knowledge of HTTP or databases.
* **Repositories** (`internal/repository/`): Data access abstractions that implement interfaces defined in the use cases.
* **Controllers** (`internal/controller/`): HTTP request handling, input validation, and delegation to use cases.

### 2. Dependency Injection

The application uses constructor-based dependency injection in `internal/app/app.go` to decouple components: repositories are injected into use cases, and use cases are injected into controllers. This makes each layer independently testable.

### 3. Interface-Based Design

Business logic depends on interfaces, not concrete implementations, allowing for easy mocking in tests and swapping of implementations (e.g., changing the database).

```go
// internal/usecase/type.go
type AlbumRepository interface {
    GetAlbums() ([]entity.Album, error)
    // ... other methods
}

// internal/usecase/function.go
func (uc *AlbumUseCase) GetAlbums() ([]entity.Album, error) {
    return uc.repo.GetAlbums() // Works with any AlbumRepository implementation
}
```

## Testing Strategy

The clean architecture makes testing straightforward:

* **Unit Tests:** Use cases are tested with mocked repositories to isolate business logic.
* **Integration Tests:** Repositories are tested against a real test database to verify data persistence.
* **HTTP Tests:** Controllers can be tested with a mocked use case layer.

To run all tests:

```bash
go test ./... -v
```

---

**Happy Learning! 🚀**
