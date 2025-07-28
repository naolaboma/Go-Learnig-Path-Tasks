package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	domain "task-manager/Domain"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// --- Mock UseCases ---
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
