# Unit Testing for Task Management API

This document outlines the testing strategy for the Task Management API, focusing on the use of the `testify` library for comprehensive unit tests.

## 1. Objective

The primary goal is to ensure the correctness and stability of individual components of the application through isolated unit tests. This is achieved by testing domain models, use cases (business logic), and controllers (API endpoints).

## 2. Tools and Libraries

- **Go (Golang)**: The core programming language.
- **Testify**: A popular testing toolkit for Go, used for:
  - `suite`: Organizing tests into logical groups with setup and teardown methods.
  - `mock`: Creating mock objects for dependencies (e.g., repositories).
  - `assert`: Making assertions about test outcomes.

## 3. Testing Strategy

Our strategy is to test each layer of the application in isolation:

### Use Case Testing

- **Location**: `Usecases/*_test.go`
- **Approach**: The business logic resides in the use cases. We test this layer by mocking the `Repository` interfaces (`TaskRepository`, `UserRepository`). This allows us to test logic like task creation, user registration, and validation without needing a live database.
- **Example**: `TestCreateTask_Success` in `Usecases/task_usecases_test.go` verifies that the `CreateTask` use case correctly calls the repository's `Create` method.

### Controller Testing

- **Location**: `Delivery/controllers/*_test.go`
- **Approach**: Controllers handle HTTP request/response logic. We test this layer by mocking the `UseCase` interfaces. We use Go's `net/http/httptest` package to simulate HTTP requests and record responses.
- **Example**: `TestCreateTask_Success` in `Delivery/controllers/controller_test.go` checks if the `/tasks` endpoint returns a `201 Created` status code when provided with valid data, ensuring the controller correctly interacts with its use case mock.

### Edge Case Testing

We write tests for various scenarios, including:

- **Success Cases**: Verifying correct behavior with valid input.
- **Failure Cases**: Testing how the system handles errors (e.g., invalid input, database errors).
- **Authorization**: Ensuring that endpoints requiring authentication or specific roles (e.g., admin) are properly protected.

## 4. How to Run Tests

### Prerequisites

Ensure you have Go installed on your system.

### Install Dependencies

Navigate to the project root and run:

```sh
go mod tidy
```
