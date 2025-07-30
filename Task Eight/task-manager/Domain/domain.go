package domain

import (
	"errors"
	"time"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

// Custom error types for better error handling
var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrNotFound           = errors.New("not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrDuplicateEntry     = errors.New("duplicate entry")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	DueDate     time.Time `bson:"due_date" json:"due_date"`
	Status      string    `bson:"status" json:"status"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"-"`
	Role     Role   `bson:"role" json:"role"`
}

// Validation methods
func (t *Task) Validate() error {
	if t.Title == "" {
		return ErrInvalidInput
	}
	if t.DueDate.Before(time.Now()) {
		return ErrInvalidInput
	}
	return nil
}

func (u *User) Validate() error {
	if u.Username == "" || len(u.Username) < 3 {
		return ErrInvalidInput
	}
	if u.Password == "" || len(u.Password) < 6 {
		return ErrInvalidInput
	}
	return nil
}

// --- Service Interfaces ---
type IPasswordService interface {
	Hash(password string) (string, error)
	Check(password, hash string) bool
}

type IAuthService interface {
	GenerateToken(user *User) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
}

// --- Repository Interfaces ---
type ITaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id string) (*Task, error)
	Create(task Task) (*Task, error)
	Update(id string, task Task) (*Task, error)
	Delete(id string) error
}

type IUserRepository interface {
	Create(user User) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByID(id string) (*User, error)
	Promote(username string) error
	Exists(username string) (bool, error)
}

// --- UseCase Interfaces ---
type ITaskUseCase interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id string, task Task) (*Task, error)
	DeleteTask(id string) error
}

type IUserUseCase interface {
	Register(user User) (*User, error)
	Login(username, password string) (string, error)
	PromoteUser(username string, promoterID string) error
}
