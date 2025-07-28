package usecases

import (
	"errors"
	domain "task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskUseCaseTestSuite struct {
	suite.Suite
	mockRepo  *mocks.MockTaskRepository
	useCase   domain.TaskUseCase
	dummyTask domain.Task
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.MockTaskRepository)
	suite.useCase = NewTaskUseCase(suite.mockRepo)
	suite.dummyTask = domain.Task{
		ID:          "1",
		Title:       "Test Task",
		Description: "A task for testing",
		Status:      "Pending",
		DueDate:     time.Now(),
	}
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_Success() {
	// Arrange
	suite.mockRepo.On("Create", suite.dummyTask).Return(&suite.dummyTask, nil)

	// Act
	createdTask, err := suite.useCase.CreateTask(suite.dummyTask)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), createdTask)
	assert.Equal(suite.T(), suite.dummyTask.Title, createdTask.Title)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_Success() {
	// Arrange
	suite.mockRepo.On("GetByID", "1").Return(&suite.dummyTask, nil)

	// Act
	task, err := suite.useCase.GetTaskByID("1")

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), task)
	assert.Equal(suite.T(), "1", task.ID)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_NotFound() {
	// Arrange
	suite.mockRepo.On("GetByID", "2").Return(nil, errors.New("not found"))

	// Act
	task, err := suite.useCase.GetTaskByID("2")

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), task)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestGetAllTasks_Success() {
	// Arrange
	tasks := []domain.Task{suite.dummyTask}
	suite.mockRepo.On("GetAll").Return(tasks, nil)

	// Act
	retrievedTasks, err := suite.useCase.GetAllTasks()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), retrievedTasks, 1)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestUpdateTask_Success() {
	// Arrange
	updatedTask := suite.dummyTask
	updatedTask.Status = "Completed"
	suite.mockRepo.On("Update", "1", updatedTask).Return(&updatedTask, nil)

	// Act
	result, err := suite.useCase.UpdateTask("1", updatedTask)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Completed", result.Status)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_Success() {
	// Arrange
	suite.mockRepo.On("Delete", "1").Return(nil)

	// Act
	err := suite.useCase.DeleteTask("1")

	// Assert
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
