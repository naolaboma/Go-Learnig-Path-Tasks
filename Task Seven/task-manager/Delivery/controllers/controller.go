// Delivery/controllers/controller.go
package controllers

import (
	"net/http"
	"task-manager/Domain"
	"task-manager/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase usecases.TaskUseCase
}

func NewTaskController(taskUseCase usecases.TaskUseCase) *TaskController {
	return &TaskController{taskUseCase: taskUseCase}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.taskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Implement other TaskController methods...

type UserController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(userUseCase usecases.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registeredUser, err := uc.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       registeredUser.ID,
		"username": registeredUser.Username,
		"role":     registeredUser.Role,
	})
}

// Implement other UserController methods...
