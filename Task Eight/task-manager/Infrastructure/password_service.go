package infrastructure

import (
	domain "task-manager/Domain"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() domain.IPasswordService {
	return &PasswordService{}
}

func (s *PasswordService) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *PasswordService) Check(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
