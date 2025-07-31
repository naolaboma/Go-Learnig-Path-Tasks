package mocks

import (
	domain "task-manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}
func (m *MockPasswordService) Check(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) GenerateToken(user *domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) ValidateToken(tokenString string) (*domain.Claims, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Claims), args.Error(1)
}
