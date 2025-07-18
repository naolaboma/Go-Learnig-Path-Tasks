# Task Management API

This is a simple Task Management API with JWT authentication and role-based access control (Admin/User).

---

## API Documentation

### Authentication

#### Register a new user

POST /register

Request body:

```json
{
  "username": "string",
  "password": "string"
}
```

Response:

```json
{
  "id": "string",
  "username": "string",
  "role": "string"
}
```

---

#### Login

POST /login

Request body:

```json
{
  "username": "string",
  "password": "string"
}
```

Response:

```json
{
  "token": "string",
  "id": "string",
  "username": "string",
  "role": "string"
}
```

---

### Tasks

#### Get all tasks (Public)

GET /tasks

Response:

```json
[
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string",
    "created_at": "string",
    "updated_at": "string"
  }
]
```

---

#### Get task by ID (Public)

GET /tasks/\:id

Response:

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string",
  "status": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

---

#### Create task (Admin only)

POST /tasks

Headers:

```
Authorization: Bearer <token>
```

Request body:

```json
{
  "title": "string",
  "description": "string",
  "due_date": "string", // e.g. "2025-09-01T17:00:00Z"
  "status": "string"
}
```

Response:

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string", // e.g. "2025-09-01T17:00:00Z"
  "status": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

---

#### Update task (Admin only)

PUT /tasks/\:id

Headers:

```
Authorization: Bearer <token>
```

Request body:

```json
{
  "title": "string",
  "description": "string",
  "due_date": "string", // e.g. "2025-09-01T17:00:00Z"
  "status": "string"
}
```

Response:

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "due_date": "string", // e.g. "2025-09-01T17:00:00Z"
  "status": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

---

#### Delete task (Admin only)

DELETE /tasks/\:id

Headers:

```
Authorization: Bearer <token>
```

Response:

```json
{
  "message": "task deleted successfully"
}
```

---

### Admin

#### Promote user to admin (Admin only)

POST /promote

Headers:

```
Authorization: Bearer <token>
```

Request body:

```json
{
  "username": "string"
}
```

Response:

```json
{
  "message": "User promoted to admin successfully"
}
```

---

## Notes

- All admin routes require a valid JWT token with admin privileges.
- The first registered user automatically becomes an admin.
- Subsequent users are regular users unless promoted by an admin.

---

## Features

1. Preserves all your existing code
2. Adds JWT authentication
3. Implements role-based authorization (admin/user)
4. Follows the required folder structure
5. Includes all the required endpoints
6. Secures passwords with bcrypt hashing
7. Provides proper API documentation

---

## How to Use

1. Register a user (the first user will be admin)
2. Login to get a JWT token
3. Use the token in the Authorization header for protected routes
4. Admins can promote other users to admin using the `/promote` endpoint

---

## Example Workflow

1. POST /register → create first admin
2. POST /login → get admin JWT token
3. POST /tasks → create tasks (requires admin token)
4. POST /register → create normal user
5. POST /promote → promote normal user to admin

---

## Folder Structure

```
.
├── controllers/     # Request handlers
├── models/          # Data models
├── services/        # Business logic
├── docs/            # Documentation
└── main.go          # Entry point
```

---

## Tech Stack

- Go (Golang)
- MongoDB
- JWT Authentication
- Bcrypt for password hashing

---
