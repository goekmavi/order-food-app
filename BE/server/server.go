package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Home page")
	})

	app.Get("/foods", func(c *fiber.Ctx) error {
		type Food struct {
			Id       int
			Headline string
			Image    string
			Rating   float64
			Anchor   string
		}

		return c.JSON(fiber.Map{
			"foods": []Food{
				{
					Id:       1,
					Headline: "Classic Margherita Pizza",
					Rating:   4.6,
					Image:    "https://cdn.dummyjson.com/recipe-images/1.webp",
					Anchor:   "#",
				},
			},
		})
	})

	app.Get("/foods/slug", func(c *fiber.Ctx) error {
		return c.SendString("Food detail page")
	})

	app.Listen(":3000")
}
