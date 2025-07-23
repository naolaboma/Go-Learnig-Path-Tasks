// Delivery/routers/router.go
package routers

import (
	"task-manager/Delivery/controllers"
	"task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.GET("/tasks", taskController.GetAllTasks)
	r.GET("/tasks/:id", taskController.GetTaskByID)

	// Protected routes
	authRoutes := r.Group("/")
	authRoutes.Use(infrastructure.AuthMiddleware())
	{
		authRoutes.POST("/tasks", taskController.CreateTask)
		authRoutes.PUT("/tasks/:id", taskController.UpdateTask)
		authRoutes.DELETE("/tasks/:id", taskController.DeleteTask)
		authRoutes.POST("/promote", userController.PromoteUser)
	}
	return r
}