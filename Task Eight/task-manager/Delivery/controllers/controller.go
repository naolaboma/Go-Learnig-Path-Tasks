package controllers

import (
	"net/http"
	"task-manager/Delivery/dto"
	domain "task-manager/Domain"

	"github.com/gin-gonic/gin"
)

// --- USER CONTROLLER ---
type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}
func (uc *UserController) Register(c *gin.Context) {
	var req dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := domain.User{Username: req.Username, Password: req.Password}

	registeredUser, err := uc.userUseCase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := dto.UserResponse{ID: registeredUser.ID, Username: registeredUser.Username, Role: string(registeredUser.Role)}
	c.JSON(http.StatusCreated, res)
}
func (uc *UserController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	token, err := uc.userUseCase.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	res := dto.LoginResponse{Token: token}
	c.JSON(http.StatusOK, res)
}
func (uc *UserController) PromoteUser(c *gin.Context) {
	var req dto.PromoteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	promoterID, _ := c.Get("userID")
	err := uc.userUseCase.PromoteUser(req.Username, promoterID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted successfully"})
}

// --- TASK CONTROLLER ---

type TaskController struct {
	taskUseCase domain.TaskUseCase
}

func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController {
	return &TaskController{taskUseCase: taskUseCase}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var req dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := domain.Task{Title: req.Title, Description: req.Description, DueDate: req.DueDate, Status: req.Status}

	createdTask, err := tc.taskUseCase.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	res := dto.TaskResponse{ID: createdTask.ID, Title: createdTask.Title, Description: createdTask.Description, DueDate: createdTask.DueDate, Status: createdTask.Status, CreatedAt: createdTask.CreatedAt, UpdatedAt: createdTask.UpdatedAt}
	c.JSON(http.StatusCreated, res)
}
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.taskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}
	var taskResponses []dto.TaskResponse
	for _, t := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponse{ID: t.ID, Title: t.Title, Description: t.Description, DueDate: t.DueDate, Status: t.Status, CreatedAt: t.CreatedAt, UpdatedAt: t.UpdatedAt})
	}
	c.JSON(http.StatusOK, taskResponses)
}
func (tc *TaskController) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	task, err := tc.taskUseCase.GetTaskByID(taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	res := dto.TaskResponse{ID: task.ID, Title: task.Title, Description: task.Description, DueDate: task.DueDate, Status: task.Status, CreatedAt: task.CreatedAt, UpdatedAt: task.UpdatedAt}
	c.JSON(http.StatusOK, res)
}
func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var req dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := domain.Task{Title: req.Title, Description: req.Description, DueDate: req.DueDate, Status: req.Status}
	updatedTask, err := tc.taskUseCase.UpdateTask(taskID, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := dto.TaskResponse{ID: updatedTask.ID, Title: updatedTask.Title, Description: updatedTask.Description, DueDate: updatedTask.DueDate, Status: updatedTask.Status, CreatedAt: updatedTask.CreatedAt, UpdatedAt: updatedTask.UpdatedAt}
	c.JSON(http.StatusOK, res)
}
func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	err := tc.taskUseCase.DeleteTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
