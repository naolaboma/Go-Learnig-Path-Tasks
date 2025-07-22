// Usecases/task_usecases.go
package usecases

import "task-manager/Domain"

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) GetAllTasks() ([]domain.Task, error) {
	return uc.taskRepo.GetAll()
}

// Implement other TaskUseCase methods...
