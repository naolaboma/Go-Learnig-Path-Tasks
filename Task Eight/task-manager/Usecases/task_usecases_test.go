package usecases

import (
	"errors"
	domain "task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TaskUseCaseTestSuite struct {
	suite.Suite
	mockRepo  *mocks.MockTaskRepository
	useCase   domain.ITaskUseCase
	dummyTask domain.Task
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockTaskRepository)
	suite.useCase = NewTaskUseCase(suite.mockRepo)
	suite.dummyTask = domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "A task for testing",
		Status:      "pending",
		DueDate:     time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_Success() {
	suite.mockRepo.On("Create", mock.AnythingOfType("domain.Task")).Return(&suite.dummyTask, nil)
	_, err := suite.useCase.CreateTask(suite.dummyTask)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_ValidationError_EmptyTitle() {
	invalidTask := suite.dummyTask
	invalidTask.Title = ""
	_, err := suite.useCase.CreateTask(invalidTask)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_ValidationError_PastDueDate() {
	invalidTask := suite.dummyTask
	invalidTask.DueDate = time.Now().Add(-24 * time.Hour)
	_, err := suite.useCase.CreateTask(invalidTask)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_RepositoryError() {
	suite.mockRepo.On("Create", mock.AnythingOfType("domain.Task")).Return(nil, errors.New("database error"))
	_, err := suite.useCase.CreateTask(suite.dummyTask)
	assert.Error(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_Success() {
	suite.mockRepo.On("GetByID", "1").Return(&suite.dummyTask, nil)
	task, err := suite.useCase.GetTaskByID("1")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), task)
	assert.Equal(suite.T(), "1", task.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_EmptyID() {
	_, err := suite.useCase.GetTaskByID("")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_NotFound() {
	suite.mockRepo.On("GetByID", "2").Return(nil, errors.New("not found"))
	_, err := suite.useCase.GetTaskByID("2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetAllTasks_Success() {
	tasks := []domain.Task{suite.dummyTask}
	suite.mockRepo.On("GetAll").Return(tasks, nil)
	retrievedTasks, err := suite.useCase.GetAllTasks()
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), retrievedTasks, 1)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetAllTasks_RepositoryError() {
	suite.mockRepo.On("GetAll").Return(nil, errors.New("database error"))
	_, err := suite.useCase.GetAllTasks()
	assert.Error(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_Success() {
	updatedTask := suite.dummyTask
	updatedTask.Status = "completed"
	suite.mockRepo.On("GetByID", "1").Return(&suite.dummyTask, nil)
	suite.mockRepo.On("Update", "1", mock.AnythingOfType("domain.Task")).Return(&updatedTask, nil)
	_, err := suite.useCase.UpdateTask("1", updatedTask)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_EmptyID() {
	_, err := suite.useCase.UpdateTask("", suite.dummyTask)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_ValidationError() {
	invalidTask := suite.dummyTask
	invalidTask.Title = ""
	_, err := suite.useCase.UpdateTask("1", invalidTask)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_NotFound() {
	suite.mockRepo.On("GetByID", "2").Return(nil, errors.New("not found"))
	_, err := suite.useCase.UpdateTask("2", suite.dummyTask)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_Success() {
	suite.mockRepo.On("GetByID", "1").Return(&suite.dummyTask, nil)
	suite.mockRepo.On("Delete", "1").Return(nil)
	err := suite.useCase.DeleteTask("1")
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_EmptyID() {
	err := suite.useCase.DeleteTask("")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrInvalidInput, err)
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_NotFound() {
	suite.mockRepo.On("GetByID", "2").Return(nil, errors.New("not found"))
	err := suite.useCase.DeleteTask("2")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.ErrNotFound, err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
