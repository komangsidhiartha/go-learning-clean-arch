# Go Clean Architecture Template

A learning-focused Go template that demonstrates clean architecture principles by combining two official Go tutorials: [Database Access](https://go.dev/doc/tutorial/database-access) and [Web Service with Gin](https://go.dev/doc/tutorial/web-service-gin).

## 🎯 Learning Objectives

This template is designed to help you learn **Go project templating and architecture patterns** rather than just Go language syntax. You'll learn:

- **Clean Architecture** implementation in Go
- **Dependency Injection** patterns
- **Project structure** best practices
- **Layer separation** (controllers, use cases, repositories)
- **Interface-based design** for testability
- **Go module organization** and package management

## 🏗️ Architecture Overview

This project follows the [evrone/go-clean-template](https://github.com/evrone/go-clean-template) approach with a simplified structure:

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

## 🚀 Getting Started

### Prerequisites

- Go 1.21+ 
- MariaDB (for database operations)

### Installation

1. **Clone the template:**
   ```bash
   git clone <your-repo-url>
   cd go-template
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the application:**
   ```bash
   go run cmd/app/main.go
   ```

4. **Run with database migrations:**
   ```bash
   go run -tags migrate cmd/app/main.go
   ```

## 📚 What You'll Learn

### 1. Clean Architecture Layers

**Entities** (`internal/entity/`)
- Pure business models with no external dependencies
- Data Transfer Objects (DTOs) for API communication

**Use Cases** (`internal/usecase/`)
- Business logic implementation
- Orchestrates between different repositories and external services
- No knowledge of HTTP, database, or external APIs

**Repositories** (`internal/repository/`)
- Data access abstraction
- Implements interfaces defined in use cases
- Can be easily swapped (SQLite → PostgreSQL → MongoDB)

**Controllers** (`internal/controller/`)
- HTTP request handling
- Input validation and response formatting
- No business logic, only delegation to use cases

### 2. Dependency Injection

The application uses constructor-based dependency injection:

```go
// internal/app/app.go
func NewApp() *App {
    // Create repositories
    albumRepo := repository.NewAlbumRepository(db)
    
    // Create use cases with injected dependencies
    albumUseCase := usecase.NewAlbumUseCase(albumRepo)
    
    // Create controllers with injected use cases
    albumHandler := handler.NewAlbumHandler(albumUseCase)
    
    return &App{
        albumHandler: albumHandler,
        // ... other dependencies
    }
}
```

### 3. Interface-Based Design

Business logic depends on interfaces, not concrete implementations:

```go
// internal/usecase/type.go
type AlbumRepository interface {
    GetAlbums() ([]entity.Album, error)
    GetAlbumByID(id string) (*entity.Album, error)
    AddAlbum(album entity.Album) error
}

// internal/usecase/function.go
func (uc *AlbumUseCase) GetAlbums() ([]entity.Album, error) {
    return uc.repo.GetAlbums() // Works with any implementation
}
```

## 🔧 Key Features

### Database Integration
- **MariaDB** for production-ready database operations
- **GORM** for database operations
- **Auto-migrations** with build tags
- **Repository pattern** for data access

### HTTP API
- **Gin framework** for routing
- **RESTful endpoints** for album management
- **Versioned API** structure (`/v1/albums`)
- **Swagger documentation** ready

### Configuration
- **Environment-based** configuration
- **Structured logging** setup
- **Graceful shutdown** handling

## 📖 Tutorial Features

This template combines two Go tutorials:

1. **Database Access Tutorial**
   - Album entity with CRUD operations
   - MariaDB database integration
   - Repository pattern implementation

2. **Web Service Tutorial**
   - RESTful HTTP endpoints
   - JSON request/response handling
   - Middleware integration

## 🧪 Testing Strategy

The clean architecture makes testing straightforward:

- **Unit tests** for use cases with mocked repositories
- **Integration tests** for repositories with test database
- **HTTP tests** for controllers with mocked use cases

## 🔄 Customization Guide

### Adding New Features

1. **Create Entity** in `internal/entity/`
2. **Define Repository Interface** in `internal/usecase/type.go`
3. **Implement Repository** in `internal/repository/`
4. **Add Use Case** in `internal/usecase/`
5. **Create Controller** in `internal/controller/http/v1/`
6. **Update Router** in `internal/controller/http/v1/router.go`

### Changing Database

Replace the repository implementation in `internal/repository/`:

```go
// For PostgreSQL
type AlbumRepository struct {
    db *gorm.DB
}

// For MongoDB
type AlbumRepository struct {
    collection *mongo.Collection
}

// For SQLite (development)
type AlbumRepository struct {
    db *gorm.DB
}
```

## 📁 Project Structure Deep Dive

### `cmd/app/main.go`
- Application entry point
- Configuration loading
- Graceful shutdown handling

### `internal/app/app.go`
- Dependency injection setup
- Application lifecycle management
- Server initialization

### `internal/controller/http/v1/`
- HTTP request handling
- Input validation
- Response formatting
- Swagger documentation

### `internal/usecase/`
- Business logic implementation
- Transaction management
- External service orchestration

### `internal/repository/`
- Data persistence logic
- Database query optimization
- Caching strategies

## 🎓 Learning Path

1. **Start with the structure** - Understand how layers are organized
2. **Follow the data flow** - HTTP → Controller → Use Case → Repository → Database
3. **Study dependency injection** - See how interfaces decouple components
4. **Experiment with changes** - Try adding new endpoints or changing the database
5. **Write tests** - Practice testing each layer independently

## 🤝 Contributing

This is a learning template. Feel free to:

- Add new features following the established patterns
- Improve the architecture
- Add more comprehensive examples
- Enhance documentation

## 📚 Additional Resources

- [Go Clean Architecture Template](https://github.com/evrone/go-clean-template) - Original inspiration
- [Go Database Tutorial](https://go.dev/doc/tutorial/database-access) - Database concepts
- [Go Web Service Tutorial](https://go.dev/doc/tutorial/web-service-gin) - HTTP handling
- [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Happy Learning! 🚀**

This template is designed to help you understand Go project organization and clean architecture principles. Focus on the patterns and structure rather than memorizing syntax - that's how you'll become a better Go developer!
