package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Routers - структура для хранения зависимостей роутов
type Routers struct {
	Service Service
}

// NewRouters - конструктор для настройки API
func NewRouters(r *Routers) *fiber.App {
	app := fiber.New()

	// Настройка CORS (разрешенные методы, заголовки, авторизация)
	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))

	// Инициализируем handler
	handler := NewHandler(r.Service)

	// Группа маршрутов с авторизацией
	apiGroup := app.Group("/v1")

	// Пользователи
	apiGroup.Post("/users", handler.CreateUser)

	// Задачи
	apiGroup.Post("/tasks", handler.CreateTask)
	apiGroup.Get("/tasks", handler.GetTasks)
	apiGroup.Get("/tasks/:id", handler.GetTaskById)
	apiGroup.Put("/tasks/:id/status", handler.UpdateTaskStatus)
	apiGroup.Delete("/tasks/:id", handler.DeleteTask)

	// Задачи по пользователю
	apiGroup.Get("/tasks/user/:user_id", handler.GetTasksByUserId)
	apiGroup.Get("/task/user/:user_id", handler.GetTaskByUserId)

	return app
}
