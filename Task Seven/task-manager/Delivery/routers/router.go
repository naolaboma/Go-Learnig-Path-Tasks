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

	// All task routes require authentication
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(infrastructure.AuthMiddleware())
	{
		taskRoutes.GET("/", taskController.GetAllTasks)
		taskRoutes.GET("/:id", taskController.GetTaskByID)

		// Admin-only task routes
		adminTaskRoutes := taskRoutes.Group("/")
		adminTaskRoutes.Use(infrastructure.AdminOnly())
		{
			adminTaskRoutes.POST("/", taskController.CreateTask)
			adminTaskRoutes.PUT("/:id", taskController.UpdateTask)
			adminTaskRoutes.DELETE("/:id", taskController.DeleteTask)
		}
	}

	// Admin-only user management routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(infrastructure.AuthMiddleware(), infrastructure.AdminOnly())
	{
		adminRoutes.POST("/promote", userController.PromoteUser)
	}

	return r
}
