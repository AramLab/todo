package server

import (
	"context"
	"github.com/AramLab/todo/internal/task/dto"
	"github.com/AramLab/todo/internal/task/models"
	"github.com/AramLab/todo/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=Service --output=../service/mocks --case=underscore
type Service interface {
	CreateUser(ctx context.Context, req dto.UserRequest) (int, error)
	CreateTask(ctx context.Context, req dto.TaskRequest) (string, error)
	GetTaskByID(ctx context.Context, id int) (*models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTaskByUserID(ctx context.Context, userID int) (*models.Task, error)
	GetTasksByUserID(ctx context.Context, userID int) ([]models.Task, error)

	// Переделать функцию: нужно обновлять всю таску
	UpdateTaskStatus(ctx context.Context, id string, status string) error
	DeleteTask(ctx context.Context, id int) error
}

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx *fiber.Ctx) error {
	var req dto.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := validator.Validate(ctx.Context(), &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := h.svc.CreateUser(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create user"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *Handler) CreateTask(ctx *fiber.Ctx) error {
	var req dto.TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := validator.Validate(ctx.Context(), &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := h.svc.CreateTask(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create task"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func (h *Handler) GetTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	task, err := h.svc.GetTaskByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (h *Handler) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := h.svc.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch tasks"})
	}
	return ctx.Status(fiber.StatusOK).JSON(tasks)
}

func (h *Handler) GetTaskByUserId(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userID"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userID"})
	}

	task, err := h.svc.GetTaskByUserID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "task not found for user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(task)
}

func (h *Handler) GetTasksByUserId(ctx *fiber.Ctx) error {
	userID, err := strconv.Atoi(ctx.Params("userID"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid userID"})
	}

	tasks, err := h.svc.GetTasksByUserID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get tasks for user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(tasks)
}

func (h *Handler) UpdateTaskStatus(ctx *fiber.Ctx) error {
	var req dto.TaskStatusUpdateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := validator.Validate(ctx.Context(), &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.svc.UpdateTaskStatus(ctx.Context(), req.ID, req.Status); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update task status"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "task status updated"})
}

func (h *Handler) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := h.svc.DeleteTask(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete task"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "task deleted"})
}
