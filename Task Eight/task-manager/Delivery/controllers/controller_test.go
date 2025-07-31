package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-manager/Delivery/dto"
	domain "task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ControllerTestSuite struct {
	suite.Suite
	router          *gin.Engine
	mockTaskUseCase *mocks.MockTaskUseCase
	mockUserUseCase *mocks.MockUserUseCase
	taskController  *TaskController
	userController  *UserController
}

func (suite *ControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.router = gin.New()

	suite.mockTaskUseCase = new(mocks.MockTaskUseCase)
	suite.mockUserUseCase = new(mocks.MockUserUseCase)

	suite.taskController = NewTaskController(suite.mockTaskUseCase)
	suite.userController = NewUserController(suite.mockUserUseCase)

	// Setup routes
	suite.router.POST("/register", suite.userController.Register)
	suite.router.POST("/login", suite.userController.Login)
	suite.router.GET("/tasks", suite.taskController.GetAllTasks)
	suite.router.GET("/tasks/:id", suite.taskController.GetTaskByID)
	suite.router.POST("/tasks", suite.taskController.CreateTask)
	suite.router.PUT("/tasks/:id", suite.taskController.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.taskController.DeleteTask)
}

func (suite *ControllerTestSuite) TestRegister_Success() {
	user := domain.User{
		ID:       "1",
		Username: "testuser",
		Role:     domain.RoleUser,
	}

	suite.mockUserUseCase.On("Register", mock.AnythingOfType("domain.User")).Return(&user, nil)

	reqBody := dto.RegisterUserRequest{
		Username: "testuser",
		Password: "password123",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response dto.UserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "testuser", response.Username)

	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestRegister_InvalidRequest() {
	reqBody := dto.RegisterUserRequest{
		Username: "",
		Password: "password123",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *ControllerTestSuite) TestRegister_UseCaseError() {
	suite.mockUserUseCase.On("Register", mock.AnythingOfType("domain.User")).Return(nil, domain.ErrDuplicateEntry)

	reqBody := dto.RegisterUserRequest{
		Username: "testuser",
		Password: "password123",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)

	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestLogin_Success() {
	suite.mockUserUseCase.On("Login", "testuser", "password123").Return("jwt-token", nil)

	reqBody := dto.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response dto.LoginResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "jwt-token", response.Token)

	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestLogin_InvalidCredentials() {
	suite.mockUserUseCase.On("Login", "testuser", "wrongpassword").Return("", domain.ErrInvalidCredentials)

	reqBody := dto.LoginRequest{
		Username: "testuser",
		Password: "wrongpassword",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)

	suite.mockUserUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetAllTasks_Success() {
	tasks := []domain.Task{
		{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description 1",
			Status:      "pending",
			DueDate:     time.Now().Add(24 * time.Hour),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	suite.mockTaskUseCase.On("GetAllTasks").Return(tasks, nil)

	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response []dto.TaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), response, 1)
	assert.Equal(suite.T(), "Task 1", response[0].Title)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetAllTasks_Error() {
	suite.mockTaskUseCase.On("GetAllTasks").Return(nil, domain.ErrNotFound)

	req := httptest.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetTaskByID_Success() {
	task := domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockTaskUseCase.On("GetTaskByID", "1").Return(&task, nil)

	req := httptest.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response dto.TaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Task 1", response.Title)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestGetTaskByID_NotFound() {
	suite.mockTaskUseCase.On("GetTaskByID", "999").Return(nil, domain.ErrNotFound)

	req := httptest.NewRequest("GET", "/tasks/999", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNotFound, w.Code)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestCreateTask_Success() {
	task := domain.Task{
		ID:          "1",
		Title:       "New Task",
		Description: "New Description",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.mockTaskUseCase.On("CreateTask", mock.AnythingOfType("domain.Task")).Return(&task, nil)

	reqBody := dto.CreateTaskRequest{
		Title:       "New Task",
		Description: "New Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response dto.TaskResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "New Task", response.Title)

	suite.mockTaskUseCase.AssertExpectations(suite.T())
}

func (suite *ControllerTestSuite) TestCreateTask_InvalidRequest() {
	reqBody := dto.CreateTaskRequest{
		Title:       "",
		Description: "New Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "pending",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func TestControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}
