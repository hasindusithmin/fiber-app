package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	arithmetic := fiber.New()
	arithmetic.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("golang maths")
	})
	arithmetic.Get("/add/:n1/:n2", func(c *fiber.Ctx) error {
		f1, e1 := strconv.ParseFloat(c.Params("n1"), 32)
		f2, e2 := strconv.ParseFloat(c.Params("n2"), 32)
		if e1 != nil || e2 != nil {
			return fiber.NewError(500, "path variable isn't numeric")
		}
		return c.JSON(map[string]float64{"addition": (f1 + f2)})
	})
	arithmetic.Get("/sub/:n1/:n2", func(c *fiber.Ctx) error {
		f1, e1 := strconv.ParseFloat(c.Params("n1"), 32)
		f2, e2 := strconv.ParseFloat(c.Params("n2"), 32)
		if e1 != nil || e2 != nil {
			return fiber.NewError(500, "path variable isn't numeric")
		}
		return c.JSON(map[string]float64{"subtraction": (f1 - f2)})
	})
	arithmetic.Get("/mul/:n1/:n2", func(c *fiber.Ctx) error {
		f1, e1 := strconv.ParseFloat(c.Params("n1"), 32)
		f2, e2 := strconv.ParseFloat(c.Params("n2"), 32)
		if e1 != nil || e2 != nil {
			return fiber.NewError(500, "path variable isn't numeric")
		}
		return c.JSON(map[string]float64{"multiplication": (f1 * f2)})
	})
	arithmetic.Get("/div/:n1/:n2", func(c *fiber.Ctx) error {
		f1, e1 := strconv.ParseFloat(c.Params("n1"), 32)
		f2, e2 := strconv.ParseFloat(c.Params("n2"), 32)
		if e1 != nil || e2 != nil {
			return fiber.NewError(500, "path variable isn't numeric")
		}
		return c.JSON(map[string]float64{"division": (f1 / f2)})
	})
	app.Mount("/arithmetic", arithmetic)

	app.Listen(":3000")
}
