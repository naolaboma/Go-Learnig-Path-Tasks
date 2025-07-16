package data

import (
	"errors"
	"task_manager/models"
	"time"
)

var tasks []models.Task
var nextID = 1

func GetAllTasks() []models.Task {
	return tasks
}
func GetTaskByID(id int) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
func CreateTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	tasks = append(tasks, task)
	return task
}
func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			updatedTask.ID = id
			updatedTask.CreatedAt = tasks[i].CreatedAt
			updatedTask.UpdatedAt = time.Now()
			tasks[i] = updatedTask
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
func DeleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
