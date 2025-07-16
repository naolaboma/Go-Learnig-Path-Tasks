# Task Manager API Documentation

## Base URL

`http://localhost:8080`

## Endpoints

### 1. Get All Tasks

**GET /tasks**`

- **Response:** `200 OK`

```json
[
  {
    "id": 1,
    "title": "Sample Task",
    "description": "This is a sample",
    "due_date": "2025-07-30",
    "status": "pending"
  }
]
2. Get Task by ID
GET /tasks/:id

Response: 200 OK

404 Not Found if task doesn’t exist.

3. Create Task
POST /tasks
{
  "title": "New Task",
  "description": "Some details",
  "due_date": "2025-08-01",
  "status": "pending"
}
Response: 201 Created

4. Update Task
PUT /tasks/:id
{
  "title": "Updated Task",
  "description": "Updated details",
  "due_date": "2025-08-05",
  "status": "completed"
}
Response: 200 OK

404 Not Found if task doesn’t exist.

5. Delete Task
DELETE /tasks/:id

Response: 200 OK

404 Not Found if task doesn’t exist.
```
