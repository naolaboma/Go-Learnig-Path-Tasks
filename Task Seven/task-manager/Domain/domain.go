package domain

import (
	"time"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Task struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     Role   `json:"role" bson:"role"`
}

// Repository Interfaces
type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task Task) (*Task, error)
	Update(id string, task Task) (*Task, error)
	Delete(id string) error
}

type UserRepository interface {
	Register(user User) (*User, error)
	Login(username, password string) (*User, error)
	PromoteUser(username string, promoterID string) error
	GetByID(id string) (*User, error)
}

// Use Case Interfaces
type TaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id string, task Task) (*Task, error)
	DeleteTask(id string) error
}

type UserUseCase interface {
	Register(user User) (*User, error)
	Login(username, password string) (*User, error)
	PromoteUser(username string, promoterID string) error
}
