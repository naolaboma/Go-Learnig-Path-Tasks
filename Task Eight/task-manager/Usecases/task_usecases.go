package usecases

import (
	domain "task-manager/Domain"
	"time"
)

type TaskUseCase struct {
	taskRepo domain.ITaskRepository
}

func NewTaskUseCase(taskRepo domain.ITaskRepository) domain.ITaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) GetAllTasks() ([]domain.Task, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (uc *TaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	return task, nil
}

func (uc *TaskUseCase) CreateTask(task domain.Task) (*domain.Task, error) {
	// Validate task
	if err := task.Validate(); err != nil {
		return nil, err
	}

	// Set default status if not provided
	if task.Status == "" {
		task.Status = "pending"
	}

	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	createdTask, err := uc.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (uc *TaskUseCase) UpdateTask(id string, task domain.Task) (*domain.Task, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	// Validate task
	if err := task.Validate(); err != nil {
		return nil, err
	}

	// Check if task exists
	existingTask, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	// Preserve original creation time
	task.CreatedAt = existingTask.CreatedAt
	task.UpdatedAt = time.Now()

	updatedTask, err := uc.taskRepo.Update(id, task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	if id == "" {
		return domain.ErrInvalidInput
	}

	// Check if task exists
	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return uc.taskRepo.Delete(id)
}
