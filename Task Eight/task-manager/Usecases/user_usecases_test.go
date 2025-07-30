package usecases

import (
	"errors"
	domain "task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockUserRepo    *mocks.MockUserRepository
	mockPasswordSvc *mocks.MockPasswordService
	mockAuthSvc     *mocks.MockAuthService
	useCase         domain.IUserUseCase
	dummyUser       domain.User
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.MockUserRepository)
	suite.mockPasswordSvc = new(mocks.MockPasswordService)
	suite.mockAuthSvc = new(mocks.MockAuthService)
	suite.useCase = NewUserUseCase(suite.mockUserRepo, suite.mockPasswordSvc, suite.mockAuthSvc)
	suite.dummyUser = domain.User{
		ID:       "1",
		Username: "testuser",
		Password: "hashedpassword",
		Role:     domain.RoleUser,
	}
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	suite.mockUserRepo.On("Exists", "testuser").Return(false, nil)
	suite.mockPasswordSvc.On("Hash", "password123").Return("hashedpassword", nil)
	suite.mockUserRepo.On("Create", mock.AnythingOfType("domain.User")).Return(&suite.dummyUser, nil)

	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	result, err := suite.useCase.Register(user)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), "testuser", result.Username)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordSvc.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegister_ValidationError_ShortUsername() {
	user := domain.User{
		Username: "ab",
		Password: "password123",
	}

	_, err := suite.useCase.Register(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *UserUseCaseTestSuite) TestRegister_ValidationError_ShortPassword() {
	user := domain.User{
		Username: "testuser",
		Password: "123",
	}

	_, err := suite.useCase.Register(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *UserUseCaseTestSuite) TestRegister_UserAlreadyExists() {
	suite.mockUserRepo.On("Exists", "testuser").Return(true, nil)

	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	_, err := suite.useCase.Register(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrDuplicateEntry, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegister_PasswordHashError() {
	suite.mockUserRepo.On("Exists", "testuser").Return(false, nil)
	suite.mockPasswordSvc.On("Hash", "password123").Return("", errors.New("hash error"))

	user := domain.User{
		Username: "testuser",
		Password: "password123",
	}

	_, err := suite.useCase.Register(user)
	assert.Error(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordSvc.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	suite.mockUserRepo.On("GetByUsername", "testuser").Return(&suite.dummyUser, nil)
	suite.mockPasswordSvc.On("Check", "password123", mock.AnythingOfType("string")).Return(true)
	suite.mockAuthSvc.On("GenerateToken", &suite.dummyUser).Return("jwt-token", nil)

	token, err := suite.useCase.Login("testuser", "password123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "jwt-token", token)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordSvc.AssertExpectations(suite.T())
	suite.mockAuthSvc.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_EmptyCredentials() {
	_, err := suite.useCase.Login("", "password123")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)

	_, err = suite.useCase.Login("testuser", "")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *UserUseCaseTestSuite) TestLogin_UserNotFound() {
	suite.mockUserRepo.On("GetByUsername", "nonexistent").Return(nil, errors.New("not found"))

	_, err := suite.useCase.Login("nonexistent", "password123")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidCredentials, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_InvalidPassword() {
	suite.mockUserRepo.On("GetByUsername", "testuser").Return(&suite.dummyUser, nil)
	suite.mockPasswordSvc.On("Check", "wrongpassword", mock.AnythingOfType("string")).Return(false)

	_, err := suite.useCase.Login("testuser", "wrongpassword")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidCredentials, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordSvc.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_TokenGenerationError() {
	suite.mockUserRepo.On("GetByUsername", "testuser").Return(&suite.dummyUser, nil)
	suite.mockPasswordSvc.On("Check", "password123", mock.AnythingOfType("string")).Return(true)
	suite.mockAuthSvc.On("GenerateToken", &suite.dummyUser).Return("", errors.New("token error"))

	_, err := suite.useCase.Login("testuser", "password123")
	assert.Error(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordSvc.AssertExpectations(suite.T())
	suite.mockAuthSvc.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_Success() {
	promoter := domain.User{
		ID:       "2",
		Username: "admin",
		Role:     domain.RoleAdmin,
	}
	userToPromote := domain.User{
		ID:       "1",
		Username: "testuser",
		Role:     domain.RoleUser,
	}

	suite.mockUserRepo.On("GetByID", "2").Return(&promoter, nil)
	suite.mockUserRepo.On("GetByUsername", "testuser").Return(&userToPromote, nil)
	suite.mockUserRepo.On("Promote", "testuser").Return(nil)

	err := suite.useCase.PromoteUser("testuser", "2")
	assert.NoError(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_EmptyInputs() {
	err := suite.useCase.PromoteUser("", "2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)

	err = suite.useCase.PromoteUser("testuser", "")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_PromoterNotFound() {
	suite.mockUserRepo.On("GetByID", "2").Return(nil, errors.New("not found"))

	err := suite.useCase.PromoteUser("testuser", "2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_PromoterNotAdmin() {
	promoter := domain.User{
		ID:       "2",
		Username: "user",
		Role:     domain.RoleUser,
	}

	suite.mockUserRepo.On("GetByID", "2").Return(&promoter, nil)

	err := suite.useCase.PromoteUser("testuser", "2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrForbidden, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_UserToPromoteNotFound() {
	promoter := domain.User{
		ID:       "2",
		Username: "admin",
		Role:     domain.RoleAdmin,
	}

	suite.mockUserRepo.On("GetByID", "2").Return(&promoter, nil)
	suite.mockUserRepo.On("GetByUsername", "nonexistent").Return(nil, errors.New("not found"))

	err := suite.useCase.PromoteUser("nonexistent", "2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_UserAlreadyAdmin() {
	promoter := domain.User{
		ID:       "2",
		Username: "admin",
		Role:     domain.RoleAdmin,
	}
	userToPromote := domain.User{
		ID:       "1",
		Username: "testuser",
		Role:     domain.RoleAdmin,
	}

	suite.mockUserRepo.On("GetByID", "2").Return(&promoter, nil)
	suite.mockUserRepo.On("GetByUsername", "testuser").Return(&userToPromote, nil)

	err := suite.useCase.PromoteUser("testuser", "2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
