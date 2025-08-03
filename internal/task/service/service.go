package service

import (
	"context"
	"github.com/AramLab/todo/internal/task/dto"
	"github.com/AramLab/todo/internal/task/models"
	repo "github.com/AramLab/todo/internal/task/repository"
	"github.com/AramLab/todo/internal/task/server"

	"go.uber.org/zap"
)

type service struct {
	repo repo.Repository
	log  *zap.SugaredLogger
}

func NewService(r repo.Repository, log *zap.SugaredLogger) server.Service {
	return &service{repo: r, log: log}
}

func (s *service) CreateUser(ctx context.Context, req dto.UserRequest) (int, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	id, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		s.log.Errorf("CreateUser: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *service) CreateTask(ctx context.Context, req dto.TaskRequest) (int, error) {
	task := models.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	id, err := s.repo.CreateTask(ctx, task)
	if err != nil {
		s.log.Errorf("CreateTask: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *service) GetTaskByID(ctx context.Context, id int) (*models.Task, error) {
	task, err := s.repo.GetTaskByID(ctx, id)
	if err != nil {
		s.log.Errorf("GetTaskById: %v", err)
		return nil, err
	}
	return task, nil
}

func (s *service) GetTasks(ctx context.Context) ([]models.Task, error) {
	tasks, err := s.repo.GetTasks(ctx)
	if err != nil {
		s.log.Errorf("GetTasks: %v", err)
		return nil, err
	}
	return tasks, nil
}

func (s *service) GetTaskByUserID(ctx context.Context, userID int) (*models.Task, error) {
	task, err := s.repo.GetTaskByUserID(ctx, userID)
	if err != nil {
		s.log.Errorf("GetTaskByUserId: %v", err)
		return nil, err
	}
	return task, nil
}

func (s *service) GetTasksByUserID(ctx context.Context, userID int) ([]models.Task, error) {
	tasks, err := s.repo.GetTasksByUserID(ctx, userID)
	if err != nil {
		s.log.Errorf("GetTasksByUserId: %v", err)
		return nil, err
	}
	return tasks, nil
}

func (s *service) UpdateTaskStatus(ctx context.Context, id int, status string) error {
	if err := s.repo.UpdateTaskStatus(ctx, id, status); err != nil {
		s.log.Errorf("UpdateTaskStatus: %v", err)
		return err
	}
	return nil
}

func (s *service) DeleteTask(ctx context.Context, id int) error {
	if err := s.repo.DeleteTask(ctx, id); err != nil {
		s.log.Errorf("DeleteTask: %v", err)
		return err
	}
	return nil
}
