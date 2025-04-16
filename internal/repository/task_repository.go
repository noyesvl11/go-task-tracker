package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (s TaskRepositoryImpl) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := s.db.Find(&tasks).Error
	return tasks, err
}

func (s TaskRepositoryImpl) GetById(id int) (*models.Task, error) {
	var task models.Task
	err := s.db.First(&task, id).Error
	return &task, err
}

func (s TaskRepositoryImpl) Create(task *models.Task) error {
	return s.db.Create(task).Error
}

func (s TaskRepositoryImpl) Update(id int, task *models.TaskEdit) error {
	return s.db.Model(&models.Task{}).Where("id = ?", id).Omit("id, CreatedAt").Updates(task).Error
}

func (s TaskRepositoryImpl) Delete(taskID int) error {
	return s.db.Delete(&models.Task{}, taskID).Error
}
