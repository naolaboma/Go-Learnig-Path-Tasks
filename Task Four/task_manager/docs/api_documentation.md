### Task Manager API Documentation

## Base URL

`http://localhost:8080

# Task Manager API

## 1️⃣ Get All Tasks

GET /tasks

Response: 200 OK
[
{
"id": 1,
"title": "Sample Task",
"description": "This is a sample",
"due_date": "2025-07-30",
"status": "pending"
}
]

---

## 2️⃣ Get Task by ID

GET /tasks/:id

Response:
200 OK # Returns the task details
404 Not Found # If the task doesn’t exist

---

## 3️⃣ Create Task

POST /tasks

Request Body:
{
"title": "New Task",
"description": "Some details",
"due_date": "2025-08-01",
"status": "pending"
}

Response:
201 Created # Returns the created task

---

## 4️⃣ Update Task

PUT /tasks/:id

Request Body:
{
"title": "Updated Task",
"description": "Updated details",
"due_date": "2025-08-05",
"status": "completed"
}

Response:
200 OK # Returns the updated task
404 Not Found # If the task doesn’t exist

---

## 5️⃣ Delete Task

DELETE /tasks/:id

Response:
200 OK # If deleted successfully
404 Not Found # If the task doesn’t exist
