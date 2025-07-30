# Task Manager API - Clean Architecture Implementation

A robust task management API built with Go, following clean architecture principles with comprehensive testing.

## 🏗️ Architecture Overview

This application follows Clean Architecture principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                    Delivery Layer                          │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │
│  │ Controllers │  │   Routers   │  │   Middleware    │   │
│  └─────────────┘  └─────────────┘  └─────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                    Use Cases Layer                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │
│  │Task UseCase │  │User UseCase │  │ Business Logic  │   │
│  └─────────────┘  └─────────────┘  └─────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                   Domain Layer                             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │
│  │   Entities  │  │  Interfaces │  │ Business Rules  │   │
│  └─────────────┘  └─────────────┘  └─────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                Infrastructure Layer                        │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────┐   │
│  │ Repositories│  │   Services  │  │   External      │   │
│  │             │  │             │  │   Dependencies  │   │
│  └─────────────┘  └─────────────┘  └─────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## 🚀 Features

### Core Functionality

- **User Management**: Registration, authentication, role-based access
- **Task Management**: CRUD operations with validation
- **Authentication**: JWT-based authentication
- **Authorization**: Role-based access control (Admin/User)

### Technical Features

- **Clean Architecture**: Clear separation of concerns
- **Dependency Injection**: Proper dependency management
- **Configuration Management**: Environment-based configuration
- **Comprehensive Testing**: Unit, integration, and E2E tests
- **Error Handling**: Consistent error types and responses
- **Input Validation**: Domain-level validation rules

## 📁 Project Structure

```
task-manager/
├── config/                 # Configuration management
│   └── config.go
├── Delivery/              # HTTP layer
│   ├── controllers/       # HTTP request handlers
│   ├── dto/              # Data Transfer Objects
│   ├── routers/          # Route definitions
│   └── main.go           # Application entry point
├── Domain/               # Business logic layer
│   └── domain.go         # Entities, interfaces, business rules
├── Infrastructure/       # External dependencies
│   ├── auth_middleware.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/         # Data access layer
│   ├── mocks/           # Mock implementations for testing
│   ├── task_repository.go
│   └── user_repository.go
├── Usecases/            # Business logic implementation
│   ├── task_usecases.go
│   ├── task_usecases_test.go
│   ├── user_usecases.go
│   └── user_usecases_test.go
├── tests/               # Test files
├── go.mod
├── go.sum
├── README.md
└── TESTING.md
```

## 🛠️ Installation & Setup

### Prerequisites

- Go 1.21 or higher
- MongoDB 4.4 or higher
- Git

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd task-manager
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables**

   ```bash
   export MONGO_URI="mongodb://localhost:27017"
   export MONGO_DATABASE="task_manager_db"
   export JWT_SECRET="your-secret-key"
   export SERVER_PORT="8080"
   export SERVER_HOST="localhost"
   ```

4. **Start MongoDB**

   ```bash
   # Using Docker
   docker run -d -p 27017:27017 --name mongodb mongo:latest

   # Or using local installation
   mongod
   ```

5. **Run the application**
   ```bash
   go run Delivery/main.go
   ```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test suites
go test ./Usecases/...
go test ./Delivery/controllers/...

# Run with verbose output
go test -v ./...
```

### Test Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage in browser
go tool cover -html=coverage.out
```

### Test Types

1. **Unit Tests**: Test business logic in isolation

   - Use case tests with mocked dependencies
   - Controller tests with mocked use cases
   - Domain validation tests

2. **Integration Tests**: Test component interactions

   - Repository tests with test database
   - Service tests with real implementations

3. **End-to-End Tests**: Test complete workflows
   - HTTP API tests
   - Authentication flow tests

## 📚 API Documentation

### Authentication Endpoints

#### Register User

```http
POST /register
Content-Type: application/json

{
  "username": "john_doe",
  "password": "secure_password"
}
```

#### Login

```http
POST /login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "secure_password"
}
```

### Task Endpoints

#### Get All Tasks (Authenticated)

```http
GET /tasks
Authorization: Bearer <jwt_token>
```

#### Get Task by ID (Authenticated)

```http
GET /tasks/{id}
Authorization: Bearer <jwt_token>
```

#### Create Task (Admin Only)

```http
POST /tasks
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation",
  "due_date": "2024-01-15T10:00:00Z",
  "status": "pending"
}
```

#### Update Task (Admin Only)

```http
PUT /tasks/{id}
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "title": "Updated task title",
  "description": "Updated description",
  "due_date": "2024-01-20T10:00:00Z",
  "status": "in_progress"
}
```

#### Delete Task (Admin Only)

```http
DELETE /tasks/{id}
Authorization: Bearer <jwt_token>
```

### Admin Endpoints

#### Promote User (Admin Only)

```http
POST /admin/promote
Authorization: Bearer <jwt_token>
Content-Type: application/json

{
  "username": "user_to_promote"
}
```

## 🔧 Configuration

The application uses environment variables for configuration:

| Variable         | Default                     | Description               |
| ---------------- | --------------------------- | ------------------------- |
| `MONGO_URI`      | `mongodb://localhost:27017` | MongoDB connection string |
| `MONGO_DATABASE` | `task_manager_db`           | Database name             |
| `JWT_SECRET`     | `your-very-secret-key`      | JWT signing secret        |
| `SERVER_PORT`    | `8080`                      | Server port               |
| `SERVER_HOST`    | `localhost`                 | Server host               |

## 🏛️ Clean Architecture Benefits

### 1. **Separation of Concerns**

- **Domain Layer**: Contains business entities and rules
- **Use Cases**: Implement business logic
- **Controllers**: Handle HTTP requests/responses
- **Repositories**: Manage data persistence
- **Infrastructure**: Handle external dependencies

### 2. **Dependency Inversion**

- High-level modules don't depend on low-level modules
- Both depend on abstractions
- Abstractions don't depend on details

### 3. **Testability**

- Each layer can be tested in isolation
- Dependencies can be easily mocked
- Business logic is independent of infrastructure

### 4. **Maintainability**

- Clear boundaries between layers
- Easy to modify without affecting other layers
- Well-defined interfaces

## 🚀 Key Improvements Made

### 1. **Enhanced Domain Layer**

- ✅ Custom error types for better error handling
- ✅ Domain validation methods
- ✅ Clear business rules and constraints
- ✅ Proper JSON tags for serialization

### 2. **Improved Use Cases**

- ✅ Comprehensive input validation
- ✅ Business rule enforcement
- ✅ Proper error handling and propagation
- ✅ Default value handling

### 3. **Better Error Handling**

- ✅ Domain-specific error types
- ✅ Consistent error responses
- ✅ Proper HTTP status codes
- ✅ Meaningful error messages

### 4. **Configuration Management**

- ✅ Environment-based configuration
- ✅ Centralized config management
- ✅ Default values for all settings
- ✅ Easy deployment configuration

### 5. **Comprehensive Testing**

- ✅ Unit tests for all business logic
- ✅ Controller tests for HTTP layer
- ✅ Mock-based testing for isolation
- ✅ Error scenario coverage

### 6. **Security Enhancements**

- ✅ JWT token validation
- ✅ Role-based access control
- ✅ Password hashing
- ✅ Input sanitization

## 📊 Testing Strategy

### Testing Pyramid

```
    E2E Tests (Few)
        ▲
   Integration Tests (Some)
        ▲
    Unit Tests (Many)
```

### Test Coverage Goals

- **Unit Tests**: 90%+ coverage
- **Integration Tests**: Critical paths
- **E2E Tests**: Complete user workflows

### Test Categories

1. **Unit Tests**: Fast, isolated, comprehensive
2. **Integration Tests**: Component interactions
3. **End-to-End Tests**: Complete workflows
4. **Performance Tests**: Load and stress testing

## 🔍 Code Quality

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Code Coverage

```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage
go tool cover -func=coverage.out
```

## 🚀 Deployment

### Docker Deployment

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main Delivery/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### Environment Variables

```bash
# Production environment
export MONGO_URI="mongodb://production-mongo:27017"
export JWT_SECRET="your-production-secret"
export SERVER_PORT="8080"
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For support and questions:

- Create an issue in the repository
- Check the documentation
- Review the test examples

---

**Built with ❤️ using Clean Architecture principles**
