package main

import (
	"myfiber/repo"
	"myfiber/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	app.Get("/people", func(c *fiber.Ctx) error {
		return c.JSON(repo.ListPeople())
	})

	router.InitRouter(app)

	_ = app.Listen(":8080")
}