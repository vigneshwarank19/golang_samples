package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Hulk(c *fiber.Ctx) error {
	a := "hello peter"
	return c.SendString(a)
}

func spidy(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
