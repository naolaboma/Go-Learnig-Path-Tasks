package usecases

import (
	domain "task-manager/Domain"
)

type UserUseCase struct {
	userRepo        domain.IUserRepository
	passwordService domain.IPasswordService
	authService     domain.IAuthService
}

func NewUserUseCase(userRepo domain.IUserRepository, passwordService domain.IPasswordService, authService domain.IAuthService) domain.IUserUseCase {
	return &UserUseCase{
		userRepo:        userRepo,
		passwordService: passwordService,
		authService:     authService,
	}
}

func (uc *UserUseCase) Register(user domain.User) (*domain.User, error) {
	// Validate user input
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Check if user already exists
	exists, err := uc.userRepo.Exists(user.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrDuplicateEntry
	}

	// Hash password
	hashedPassword, err := uc.passwordService.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	// Set default role
	if user.Role == "" {
		user.Role = domain.RoleUser
	}

	// Set hashed password
	user.Password = hashedPassword

	// Create user
	createdUser, err := uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (uc *UserUseCase) Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", domain.ErrInvalidInput
	}

	// Get user by username
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Check password
	if !uc.passwordService.Check(password, user.Password) {
		return "", domain.ErrInvalidCredentials
	}

	// Generate token
	token, err := uc.authService.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UserUseCase) PromoteUser(username string, promoterID string) error {
	if username == "" || promoterID == "" {
		return domain.ErrInvalidInput
	}

	// Check if promoter exists and is admin
	promoter, err := uc.userRepo.GetByID(promoterID)
	if err != nil {
		return domain.ErrNotFound
	}
	if promoter.Role != domain.RoleAdmin {
		return domain.ErrForbidden
	}

	// Check if user to be promoted exists
	userToPromote, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return domain.ErrNotFound
	}

	// Prevent promoting already admin users
	if userToPromote.Role == domain.RoleAdmin {
		return domain.ErrInvalidInput
	}

	return uc.userRepo.Promote(username)
}
