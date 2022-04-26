package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s", c.Params("name"))
		return c.SendString(msg)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("%s is %s years old", c.Query("name"), c.Query("age"))
		return c.SendString(msg)
	})

	app.Get("/car/:model?", func(c *fiber.Ctx) error {
		if c.Params("model") != "" {
			return c.SendString("model " + c.Params("model"))
		}
		return c.SendStatus(200)
	})

	app.Listen(":3000")
}
