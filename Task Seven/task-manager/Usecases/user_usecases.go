// Usecases/user_usecases.go
package usecases

import "task-manager/Domain"

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) Register(user domain.User) (*domain.User, error) {
	return uc.userRepo.Register(user)
}

// Implement other UserUseCase methods...
