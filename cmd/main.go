package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stabildev/keno/handlers"
	"github.com/stabildev/keno/models"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Keno ERP",
		ServerHeader: "Keno",
	})

	models.ConnectDatabase()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Hello World!",
		})
	})

	app.Post("/api/v1/users", handlers.CreateUser)

	app.Get("/api/v1/users", handlers.GetUsers)

	app.Get("/api/v1/users/:id", handlers.GetUser)

	app.Put("/api/v1/users/:id", handlers.UpdateUser)

	app.Delete("/api/v1/users/:id", handlers.DeleteUser)

	app.Listen(":3000")
}
