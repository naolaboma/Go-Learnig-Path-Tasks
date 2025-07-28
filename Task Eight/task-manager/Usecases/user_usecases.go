package usecases

import (
	"errors"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) Register(user domain.User) (*domain.User, error) {
	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) Login(username, password string) (string, error) {
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !infrastructure.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := infrastructure.GenerateJWT(user)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func (uc *UserUseCase) PromoteUser(username string, promoterID string) error {
	promoter, err := uc.userRepo.GetByID(promoterID)
	if err != nil {
		return errors.New("promoter not found")
	}
	if promoter.Role != domain.RoleAdmin {
		return errors.New("only admins can promote users")
	}
	return uc.userRepo.Promote(username)
}
