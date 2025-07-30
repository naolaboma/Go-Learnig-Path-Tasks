package mocks

import (
	"github.com/stretchr/testify/mock"
	"task-manager/Domain"
)

// MockTaskUseCase is a mock for ITaskUseCase
type MockTaskUseCase struct {
	mock.Mock
}

func (m *MockTaskUseCase) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) CreateTask(task domain.Task) (*domain.Task, error) {
	args := m.Called(task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) UpdateTask(id string, task domain.Task) (*domain.Task, error) {
	args := m.Called(id, task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUseCase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockUserUseCase is a mock for IUserUseCase
type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Register(user domain.User) (*domain.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUseCase) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserUseCase) PromoteUser(username string, promoterID string) error {
	args := m.Called(username, promoterID)
	return args.Error(0)
} 