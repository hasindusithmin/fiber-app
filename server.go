package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
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

	app.Get("/convert/:numeric", func(c *fiber.Ctx) error {
		num := c.Params("numeric")
		int, err := strconv.ParseInt(num, 10, 0)
		if err != nil {
			return fiber.NewError(500, "path variable isn't numeric")
		}
		msg := fmt.Sprintf("number %d", int)
		return c.SendString(msg)
	})

	app.Listen(":3000")
}
