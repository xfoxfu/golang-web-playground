package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	// Listen on PORT 3000
	app.Listen(":3000")
}
