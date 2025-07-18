# Add API documentation (docs/api_documentation.md)

# Task Management API Documentation

# Authentication

## Register a new user

POST /register

Request body:
{
"username": "string",
"password": "string"
}

Response:
{
"id": "string",
"username": "string",
"role": "string"
}

## Login

POST /login

Request body:
{
"username": "string",
"password": "string"
}

Response:
{
"token": "string",
"id": "string",
"username": "string",
"role": "string"
}

# Tasks

## Get all tasks (Public)

GET /tasks

Response:
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

## Get task by ID (Public)

GET /tasks/:id

Response:
{
"id": "string",
"title": "string",
"description": "string",
"due_date": "string",
"status": "string",
"created_at": "string",
"updated_at": "string"
}

## Create task (Admin only)

POST /tasks/

Headers:
Authorization: Bearer <token>

Request body:
{
"title": "string",
"description": "string",
"due_date": "string",
"status": "string"
}

Response:
{
"id": "string",
"title": "string",
"description": "string",
"due_date": "string",
"status": "string",
"created_at": "string",
"updated_at": "string"
}

## Update task (Admin only)

PUT /tasks/:id

Headers:
Authorization: Bearer <token>

Request body:
{
"title": "string",
"description": "string",
"due_date": "string",
"status": "string"
}

Response:
{
"id": "string",
"title": "string",
"description": "string",
"due_date": "string",
"status": "string",
"created_at": "string",
"updated_at": "string"
}

## Delete task (Admin only)

DELETE /tasks/:id

Headers:
Authorization: Bearer <token>

Response:
{
"message": "task deleted successfully"
}

# Admin

## Promote user to admin (Admin only)

POST /promote

Headers:
Authorization: Bearer <token>

Request body:
{
"username": "string"
}

Response:
{
"message": "User promoted to admin successfully"
}

# Notes

- All admin routes require a valid JWT token with admin privileges
- The first registered user automatically becomes an admin
- Subsequent users are regular users unless promoted by an admin

---

# This implementation:

1. Preserves all your existing code
2. Adds JWT authentication
3. Implements role-based authorization (admin/user)
4. Follows the required folder structure
5. Includes all the required endpoints
6. Secures passwords with bcrypt hashing
7. Provides proper API documentation

# To use the API:

1. First register a user (the first user will be admin)
2. Login to get a JWT token
3. Use the token in the Authorization header for protected routes
4. Admins can promote other users to admin using the /promote endpoint
