package services

import (
	"rest-project/internal/models"
)

type TaskRepository interface {
	GetAll() ([]models.Task, error)
	GetById(id int) (*models.Task, error)
	Create(task *models.Task) error
	Update(id int, task *models.TaskEdit) error
	Delete(taskID int) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(taskRepo TaskRepository) *TaskService {
	return &TaskService{repo: taskRepo}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.repo.GetById(id)
}

func (s *TaskService) Create(title, description string, status int) (*models.Task, error) {
	task := &models.Task{
		Title:       title,
		Description: description,
		Status:      status,
	}
	err := s.repo.Create(task)
	return task, err
}

func (s *TaskService) Update(id int, taskEdit *models.TaskEdit) (*models.Task, error) {
	err := s.repo.Update(id, taskEdit)
	if err != nil {
		return nil, err
	}
	return s.GetTaskByID(id)
}

func (s *TaskService) DeleteTask(taskID int) error {
	return s.repo.Delete(taskID)
}
