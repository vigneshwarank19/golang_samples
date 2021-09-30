package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Hero(ctx *fiber.App) error {
	apigroup := ctx.Group("/api")
	apigroup.Get("/list", Hulk)
	apigroup.Get("/spidy", spidy)
	return nil
}
