package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task-manager/Delivery/dto"
	domain "task-manager/Domain"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// --- Mock MockTaskUseCases ---
type MockTaskUseCase struct {
	mock.Mock
}

// Corrected to use the 'domain' alias
func (m *MockTaskUseCase) CreateTask(task domain.Task) (*domain.Task, error) {
	args := m.Called(task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

// Corrected to use the 'domain' alias
func (m *MockTaskUseCase) GetAllTasks() ([]domain.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Task), args.Error(1)
}

// Corrected to use the 'domain' alias
func (m *MockTaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Task), args.Error(1)
}

// Corrected to use the 'domain' alias
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

// --- TaskController Test Suite ---
type TaskControllerTestSuite struct {
	suite.Suite
	router          *gin.Engine
	mockTaskUseCase *MockTaskUseCase
	taskController  *TaskController
}

func (suite *TaskControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.mockTaskUseCase = new(MockTaskUseCase)
	suite.taskController = NewTaskController(suite.mockTaskUseCase)
	suite.router.POST("/tasks", suite.taskController.CreateTask)
	suite.router.GET("/tasks/:id", suite.taskController.GetTaskByID)
}

func (suite *TaskControllerTestSuite) TestCreateTask_Success() {
	// Arrange
	// Corrected to use the 'domain' alias
	taskToCreate := domain.Task{Title: "New Task", Status: "new"}
	suite.mockTaskUseCase.On("CreateTask", mock.Anything).Return(&taskToCreate, nil)

	body, _ := json.Marshal(taskToCreate)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockTaskUseCase.AssertCalled(suite.T(), "CreateTask", mock.Anything)
}

func (suite *TaskControllerTestSuite) TestCreateTask_BindingError() {
	// Arrange
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"title":}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *TaskControllerTestSuite) TestGetTaskByID_Success() {
	// Arrange
	// Corrected to use the 'domain' alias
	expectedTask := &domain.Task{ID: "123", Title: "Found Task", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	suite.mockTaskUseCase.On("GetTaskByID", "123").Return(expectedTask, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks/123", nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	// Corrected to use the 'domain' alias
	var returnedTask domain.Task
	json.Unmarshal(w.Body.Bytes(), &returnedTask)
	assert.Equal(suite.T(), "123", returnedTask.ID)
	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *TaskControllerTestSuite) TestGetTaskByID_NotFound() {
	// Arrange
	// Corrected to use the 'domain' alias
	suite.mockTaskUseCase.On("GetTaskByID", "404").Return((*domain.Task)(nil), errors.New("task not found"))

	req, _ := http.NewRequest(http.MethodGet, "/tasks/404", nil)
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}

// --- Mock UserUseCase ---
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

// --- UserController Test Suite ---
type UserControllerTestSuite struct {
	suite.Suite
	router          *gin.Engine
	mockUserUseCase *MockUserUseCase
	userController  *UserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()
	suite.mockUserUseCase = new(MockUserUseCase)
	suite.userController = NewUserController(suite.mockUserUseCase)
	suite.router.POST("/register", suite.userController.Register)
	suite.router.POST("/login", suite.userController.Login)
}

func (suite *UserControllerTestSuite) TestRegister_Success() {
	// Arrange
	reqBody := dto.RegisterUserRequest{Username: "newuser", Password: "password"}
	resUser := &domain.User{ID: "1", Username: "newuser", Role: domain.RoleUser}

	// We use mock.AnythingOfType because the user object passed to the use case
	// inside the controller is not the same as reqBody.
	suite.mockUserUseCase.On("Register", mock.AnythingOfType("domain.User")).Return(resUser, nil)

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	var res dto.UserResponse
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(suite.T(), "newuser", res.Username)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_Success() {
	// Arrange
	reqBody := dto.LoginRequest{Username: "user", Password: "password"}
	suite.mockUserUseCase.On("Login", "user", "password").Return("mock_token", nil)

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	var res dto.LoginResponse
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(suite.T(), "mock_token", res.Token)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *UserControllerTestSuite) TestLogin_Failure() {
	// Arrange
	reqBody := dto.LoginRequest{Username: "user", Password: "wrongpassword"}
	suite.mockUserUseCase.On("Login", "user", "wrongpassword").Return("", errors.New("invalid credentials"))

	body, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	suite.router.ServeHTTP(w, req)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	suite.mockUserUseCase.AssertExpectations(suite.T())
}

// Add this function at the end of the file
func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
