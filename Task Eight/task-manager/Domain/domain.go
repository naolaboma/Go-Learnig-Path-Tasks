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
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	DueDate     time.Time `bson:"due_date"`
	Status      string    `bson:"status"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Role     Role   `bson:"role"`
}

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task Task) (*Task, error)
	Update(id string, task Task) (*Task, error)
	Delete(id string) error
}
type UserRepository interface {
	Create(user User) (*User, error)
	GetByUsername(username string) (*User, error)
	Promote(username string) error
	GetByID(id string) (*User, error)
}
type TaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id string, task Task) (*Task, error)
	DeleteTask(id string) error
}
type UserUseCase interface {
	Register(user User) (*User, error)
	Login(username, password string) (string, error)
	PromoteUser(username string, promoterID string) error
}
