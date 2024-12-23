package server

import (
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	app.Get("/foods", func(c *fiber.Ctx) error {
		return c.SendString("Foods overview page")
	})

	app.Get("/food/1", func(c *fiber.Ctx) error {
		return c.SendString("Food detail page")
	})

	app.Listen(":3000")
}
