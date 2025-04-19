package repository

import (
	"fmt"
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type CourseRepository interface {
	Create(course *models.Course) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) Create(course *models.Course) error {
	err := r.db.Create(course).Error
	if err != nil {
		fmt.Println("Error in Create repository:", err)
	}
	return err
}
