package mocks

import (
	"github.com/stretchr/testify/mock"
	"task-manager/Domain"
)

// MockUserRepository is a mock for IUserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user domain.User) (*domain.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Promote(username string) error {
	args := m.Called(username)
	return args.Error(0)
}

func (m *MockUserRepository) Exists(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}
