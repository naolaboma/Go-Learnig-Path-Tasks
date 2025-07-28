package usecases

import (
	"errors"
	"task-manager/Domain"
	"task-manager/Repositories/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		Status:      "Pending",
		DueDate:     time.Now(),
	}
}

func (suite *TaskUseCaseTestSuite) TestCreateTask_Success() {
	suite.mockRepo.On("Create", suite.dummyTask).Return(&suite.dummyTask, nil)
	_, err := suite.useCase.CreateTask(suite.dummyTask)
	assert.NoError(suite.T(), err)
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

func (suite *TaskUseCaseTestSuite) TestGetTaskByID_NotFound() {
	suite.mockRepo.On("GetByID", "2").Return(nil, errors.New("not found"))
	_, err := suite.useCase.GetTaskByID("2")
	assert.Error(suite.T(), err)
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

func (suite *TaskUseCaseTestSuite) TestUpdateTask_Success() {
	updatedTask := suite.dummyTask
	updatedTask.Status = "Completed"
	suite.mockRepo.On("Update", "1", updatedTask).Return(&updatedTask, nil)
	_, err := suite.useCase.UpdateTask("1", updatedTask)
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUseCaseTestSuite) TestDeleteTask_Success() {
	suite.mockRepo.On("Delete", "1").Return(nil)
	err := suite.useCase.DeleteTask("1")
	assert.NoError(suite.T(), err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
