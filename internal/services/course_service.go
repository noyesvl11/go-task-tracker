package services

import (
	"fmt"
	"rest-project/internal/models"
	"rest-project/internal/repository" // вот этот импорт нужен
)

type CourseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{repo}
}

func (s *CourseService) CreateCourse(course *models.Course) error {
	err := s.repo.Create(course)
	if err != nil {
		fmt.Println("Error in CreateCourse:", err)
	}
	return err
}
