package usecases

import (
	domain "task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	mockUserRepo        *mocks.MockUserRepository
	mockPasswordService *mocks.MockPasswordService
	mockAuthService     *mocks.MockAuthService
	userUseCase         domain.IUserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.mockUserRepo = new(mocks.MockUserRepository)
	// You would also initialize your new service mocks here
	// For now, we'll assume they exist for the login test
	suite.mockPasswordService = new(mocks.MockPasswordService)
	suite.mockAuthService = new(mocks.MockAuthService)
	suite.userUseCase = NewUserUseCase(suite.mockUserRepo, suite.mockPasswordService, suite.mockAuthService)
}

func (suite *UserUseCaseTestSuite) TestRegister_Success() {
	user := domain.User{Username: "test", Password: "password"}
	suite.mockUserRepo.On("Create", user).Return(&user, nil)
	_, err := suite.userUseCase.Register(user)
	assert.NoError(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin_Success() {
	username := "testuser"
	password := "password"
	hashedPassword := "hashed"
	user := &domain.User{Username: username, Password: hashedPassword}

	suite.mockUserRepo.On("GetByUsername", username).Return(user, nil)
	suite.mockPasswordService.On("Check", password, hashedPassword).Return(true)
	suite.mockAuthService.On("GenerateToken", user).Return("test_token", nil)

	token, err := suite.userUseCase.Login(username, password)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "test_token", token)
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordService.AssertExpectations(suite.T())
	suite.mockAuthService.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser_Success() {
	adminUser := &domain.User{ID: "admin1", Role: domain.RoleAdmin}
	suite.mockUserRepo.On("GetByID", "admin1").Return(adminUser, nil)
	suite.mockUserRepo.On("Promote", "testuser").Return(nil)
	err := suite.userUseCase.PromoteUser("testuser", "admin1")
	assert.NoError(suite.T(), err)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
