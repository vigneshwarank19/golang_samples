package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vigneshwaran/go_fiber/action/routes"
)

/* func rounting(app *fiber.App)  {
	app.Get("/",service.Hulk)
} */

func main() {
	app := fiber.New()
	// routes.Hero(app)
	routes.Hero(app) //path
	app.Listen(":3000")
}
