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

func (uc *TaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	return uc.taskRepo.GetByID(id)
}

func (uc *TaskUseCase) CreateTask(task domain.Task) (*domain.Task, error) {
	return uc.taskRepo.Create(task)
}

func (uc *TaskUseCase) UpdateTask(id string, task domain.Task) (*domain.Task, error) {
	return uc.taskRepo.Update(id, task)
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	return uc.taskRepo.Delete(id)
}
