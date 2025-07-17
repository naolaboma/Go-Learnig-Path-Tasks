package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", tc.GetAllTasks)
	r.GET("/tasks/:id", tc.GetTaskByID)
	r.POST("/tasks/", tc.CreateTask)
	r.PUT("/tasks/:id", tc.UpdateTask)
	r.DELETE("/tasks/:id", tc.DeleteTask)

	return r
}
