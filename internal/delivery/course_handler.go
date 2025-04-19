package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/auth"
	"rest-project/internal/models"
	"rest-project/internal/services"
)

type CourseHandler struct {
	service *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{service: service}
}

func (h *CourseHandler) CreateCourse(c *gin.Context) {
	userID, role, err := auth.ExtractUserIDAndRole(c)
	if err != nil || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 👇 Присваиваем текущего пользователя как teacher
	course.TeacherID = &userID

	if err := h.service.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create course",
			"details": err.Error(), // удобно для отладки
		})
		return
	}

	c.JSON(http.StatusCreated, course)
}
