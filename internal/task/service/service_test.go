package service

import (
	"context"
	"errors"
	"github.com/AramLab/todo/internal/task/dto"
	"github.com/AramLab/todo/internal/task/models"
	"github.com/AramLab/todo/internal/task/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"testing"
)

func TestCreateTask(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	tests := []struct {
		name        string
		taskRequest dto.TaskRequest
		expected    models.Task
		mockID      int
		mockError   error
		expectError bool
	}{
		{
			name: "task created successfully",
			taskRequest: dto.TaskRequest{
				UserID:      1,
				Title:       "Test",
				Description: "Desc",
				Status:      "todo",
			},
			expected: models.Task{
				UserID:      1,
				Title:       "Test",
				Description: "Desc",
				Status:      "todo",
			},
			mockID:      134,
			mockError:   nil,
			expectError: false,
		},
		{
			name: "repository returns error",
			taskRequest: dto.TaskRequest{
				UserID:      2,
				Title:       "Fail",
				Description: "Error task",
				Status:      "in_progress",
			},
			expected: models.Task{
				UserID:      2,
				Title:       "Fail",
				Description: "Error task",
				Status:      "in_progress",
			},
			mockID:      0,
			mockError:   errors.New("db error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			svc := NewService(mockRepo, logger)

			mockRepo.On("CreateTask", mock.Anything, tt.expected).Return(tt.mockID, tt.mockError)
			id, err := svc.CreateTask(context.Background(), tt.taskRequest)
			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, 0, id)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockID, id)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTaskByID(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	tests := []struct {
		name        string
		taskID      int
		mockTask    *models.Task
		mockError   error
		expectError bool
	}{
		{
			name:   "task got by id successfully",
			taskID: 1,
			mockTask: &models.Task{
				UserID:      1,
				Title:       "Test task",
				Description: "Desc",
				Status:      "todo",
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "task not found error",
			taskID:      2,
			mockTask:    nil,
			mockError:   errors.New("not found"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			svc := NewService(mockRepo, logger)

			mockRepo.On("GetTaskByID", mock.Anything, tt.taskID).Return(tt.mockTask, tt.mockError)
			task, err := svc.GetTaskByID(context.Background(), tt.taskID)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, nil)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.mockTask, task)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	tests := []struct {
		name        string
		taskID      int
		status      string
		mockError   error
		expectError bool
	}{
		{
			name:        "update status successfully",
			taskID:      1,
			status:      "in_progress",
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "repository returns error",
			taskID:      2,
			status:      "done",
			mockError:   errors.New("db error"),
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			svc := NewService(mockRepo, logger)

			mockRepo.On("UpdateTaskStatus", mock.Anything, tt.taskID, tt.status).Return(tt.mockError).Once()
			err := svc.UpdateTaskStatus(context.Background(), tt.taskID, tt.status)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTask(t *testing.T) {
	logger := zaptest.NewLogger(t).Sugar()

	tests := []struct {
		name        string
		taskID      int
		mockError   error
		expectError bool
	}{
		{
			name:        "delete task successfully",
			taskID:      1,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "repository returns error",
			taskID:      2,
			mockError:   errors.New("db error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			svc := NewService(mockRepo, logger)

			mockRepo.On("DeleteTask", mock.Anything, tt.taskID).Return(tt.mockError).Once()

			err := svc.DeleteTask(context.Background(), tt.taskID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
