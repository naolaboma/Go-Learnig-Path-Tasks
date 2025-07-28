package usecases

import (
	"errors"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"
	"task-manager/Repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockUserRepo *mocks.MockUserRepository
	userUseCase  domain.UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.MockUserRepository)
	suite.userUseCase = NewUserUseCase(suite.mockUserRepo)
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.mockUserRepo.On("Create", user).Return(&domain.User{ID: "1", Username: "testuser", Role: domain.RoleUser}, nil)

	registeredUser, err := suite.userUseCase.Register(user)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), registeredUser)
	assert.Equal(suite.T(), "testuser", registeredUser.Username)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestRegister_Failure() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.mockUserRepo.On("Create", user).Return(nil, errors.New("username already exists"))

	registeredUser, err := suite.userUseCase.Register(user)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), registeredUser)
	assert.Equal(suite.T(), "username already exists", err.Error())
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	password := "password"
	hashedPassword, _ := infrastructure.HashPassword(password)
	user := &domain.User{ID: "1", Username: "testuser", Password: hashedPassword, Role: domain.RoleUser}

	suite.mockUserRepo.On("GetByUsername", "testuser").Return(user, nil)

	token, err := suite.userUseCase.Login("testuser", password)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_InvalidCredentials() {
	suite.mockUserRepo.On("GetByUsername", "wronguser").Return(nil, errors.New("user not found"))

	token, err := suite.userUseCase.Login("wronguser", "password")

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), token)
	assert.Equal(suite.T(), "invalid credentials", err.Error())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_Success() {
	adminUser := &domain.User{ID: "admin1", Role: domain.RoleAdmin}
	suite.mockUserRepo.On("GetByID", "admin1").Return(adminUser, nil)
	suite.mockUserRepo.On("Promote", "testuser").Return(nil)

	err := suite.userUseCase.PromoteUser("testuser", "admin1")

	assert.NoError(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_Failure_NotAdmin() {
	regularUser := &domain.User{ID: "user1", Role: domain.RoleUser}
	suite.mockUserRepo.On("GetByID", "user1").Return(regularUser, nil)

	err := suite.userUseCase.PromoteUser("testuser", "user1")

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "only admins can promote users", err.Error())
	suite.mockUserRepo.AssertNotCalled(suite.T(), "Promote", "testuser")
}
func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
