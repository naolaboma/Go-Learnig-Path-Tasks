package domain

import (
	"time"
)

// Role defines the user roles
type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// Task represents a task in the system
type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

// User represents a user in the system
type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password"` // Use omitempty for security
	Role     Role   `json:"role" bson:"role"`
}

// TaskRepository defines the interface for task persistence
type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task Task) (*Task, error)
	Update(id string, task Task) (*Task, error)
	Delete(id string) error
}

// UserRepository defines the interface for user persistence
type UserRepository interface {
	Create(user User) (*User, error)
	GetByUsername(username string) (*User, error)
	Promote(username string) error
	GetByID(id string) (*User, error)
}

// TaskUseCase defines the business logic for tasks
type TaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id string, task Task) (*Task, error)
	DeleteTask(id string) error
}

// UserUseCase defines the business logic for users
type UserUseCase interface {
	Register(user User) (*User, error)
	Login(username, password string) (string, error) // Returns JWT token
	PromoteUser(username string, promoterID string) error
}
