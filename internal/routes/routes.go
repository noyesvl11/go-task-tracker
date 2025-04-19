package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/delivery"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
	}

	protected := r.Group("/api/v1", auth.JWTMiddleware())

	{
		courseRepo := repository.NewCourseRepository(db.DB)
		courseService := services.NewCourseService(courseRepo)
		courseHandler := delivery.NewCourseHandler(courseService)

		courseGroup := protected.Group("/courses")
		{
			courseGroup.POST("/", courseHandler.CreateCourse)
			// courseGroup.GET("/", courseHandler.GetAllCourses)
			// courseGroup.PUT("/:id", courseHandler.UpdateCourse)
			// courseGroup.DELETE("/:id", courseHandler.DeleteCourse)
		}
	}

	{
		taskRepo := repository.NewTaskRepository(db.DB)
		taskService := services.NewTaskService(taskRepo)
		taskHandler := delivery.NewTaskHandler(taskService)

		taskGroup := protected.Group("/tasks")
		{
			taskGroup.GET("/", taskHandler.GetAllTasks)
			taskGroup.GET("/:id", taskHandler.GetTask)
			taskGroup.POST("/", taskHandler.CreateTask)
			taskGroup.PUT("/:id", taskHandler.UpdateTask)
			taskGroup.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
}
