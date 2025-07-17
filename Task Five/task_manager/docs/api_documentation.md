# Task Management API Documentation

This document provides instructions for setting up and using the Task Management API, which uses MongoDB for persistent data storage.

## 1. Setup and Configuration

### Prerequisites

- Go (version 1.18 or newer)
- MongoDB (running locally or on a cloud provider like MongoDB Atlas)
- A command-line tool like `curl` or a GUI client like Postman for testing.

### Installation

1. Clone the repository to your local machine.
2. Navigate to the project directory: `cd task_manager`
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Configuration

The API connects to MongoDB using a connection URI. This URI can be configured via an environment variable named `MONGO_URI`.

If the `MONGO_URI` environment variable is not set, the application will default to `mongodb://localhost:27017`, which is suitable for a standard local MongoDB installation.

**To run the application:**

```bash
# Optional: set the environment variable if your DB is not local
export MONGO_URI="mongodb+srv://<user>:<password>@<cluster-url>/<database>?retryWrites=true&w=majority"

# Run the main application
go run main.go
```

The server will start and listen on `http://localhost:8080`.

## 2. API Endpoints

The base URL for all endpoints is `http://localhost:8080`. All task-related endpoints are grouped under the `/tasks` path.

### Task Data Model

The core resource of the API is the `Task`. It has the following structure:

```json
{
  "id": "60d5ec49f1d3c0a000123456",
  "title": "Complete Project Proposal",
  "description": "Write and submit the final proposal for the Q3 project.",
  "due_date": "2025-08-15T00:00:00Z",
  "status": "In Progress",
  "created_at": "2025-07-17T10:00:00Z",
  "updated_at": "2025-07-17T11:30:00Z"
}
```

**Note:** `id`, `created_at`, and `updated_at` are managed by the server and should not be sent when creating a new task.

---

### Create a Task

Creates a new task.

- **Endpoint:** `POST /tasks`
- **Request Body:** A JSON object representing the new task.

  **Example Request:**

  ```json
  {
    "title": "Design new API endpoints",
    "description": "Design the v2 endpoints for the user service.",
    "due_date": "2025-09-01T17:00:00Z",
    "status": "Pending"
  }
  ```

- **Success Response:** `201 Created`

  - **Body:** The newly created task object, including its server-generated `id`.

  **Example Response:**

  ```json
  {
    "id": "60d5ed8af1d3c0a000123457",
    "title": "Design new API endpoints",
    "description": "Design the v2 endpoints for the user service.",
    "due_date": "2025-09-01T17:00:00Z",
    "status": "Pending",
    "created_at": "2025-07-17T12:00:00Z",
    "updated_at": "2025-07-17T12:00:00Z"
  }
  ```

- **Error Response:** `400 Bad Request` if the JSON body is malformed.

---

### Get All Tasks

Retrieves a list of all tasks.

- **Endpoint:** `GET /tasks`
- **Success Response:** `200 OK`

  - **Body:** An array of task objects. The array will be empty if no tasks exist.

  **Example Response:**

  ```json
  [
    {
      "id": "60d5ed8af1d3c0a000123457",
      "title": "Design new API endpoints",
      "description": "Design the v2 endpoints for the user service.",
      "due_date": "2025-09-01T17:00:00Z",
      "status": "Pending",
      "created_at": "2025-07-17T12:00:00Z",
      "updated_at": "2025-07-17T12:00:00Z"
    }
  ]
  ```

---

### Get a Specific Task by ID

Retrieves a single task by its unique ID.

- **Endpoint:** `GET /tasks/:id`
- **Success Response:** `200 OK`
  - **Body:** The task object corresponding to the given ID.
- **Error Response:** `404 Not Found` if no task with the given ID exists.

---

### Update an Existing Task

Updates the details of an existing task.

- **Endpoint:** `PUT /tasks/:id`
- **Request Body:** A JSON object with the fields to be updated.

  **Example Request:**

  ```json
  {
    "title": "Design new API endpoints (v2)",
    "status": "In Progress"
  }
  ```

- **Success Response:** `200 OK`
  - **Body:** The full, updated task object.
- **Error Response:**
  - `404 Not Found` if no task with the given ID exists.
  - `400 Bad Request` if the JSON body is malformed.

---

### Delete a Task

Deletes a task by its unique ID.

- **Endpoint:** `DELETE /tasks/:id`
- **Success Response:** `200 OK`

  - **Body:** A confirmation message.

  **Example Response:**

  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

- **Error Response:** `404 Not Found` if no task with the given ID exists.
