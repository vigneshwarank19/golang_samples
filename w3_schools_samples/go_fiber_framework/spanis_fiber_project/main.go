package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vigneshwaran/spanis_project/middle"
)

func main() {
	app := fiber.New()

	app.Use(logger.New()) // fiber middleware - part-4
	app.Use(middle.Hellovk)
	// basic http part-1

	app.Get("/", func(c *fiber.Ctx) error {
		//	p := utils.CopyString(c.Path()) // - DN-"Don't know"
		/*	go func() { // asyncronose func
			time.Sleep(time.Second) //sleep time
			log.Println(p)          // path of handler
		}() //self-invok */
		t := fmt.Sprintf("%d", time.Now().Unix()) // f12
		return c.Send([]byte(t))                  //some times we can't sent datas types only send array byte can wrk
	})

	// parameter and request body - part 2

	app.Get("/name/:name", func(c *fiber.Ctx) error {
		name := c.Params("name") //parameter
		response := fmt.Sprintf("hello %s !", name)
		return c.Send([]byte(response))
	})

	// testing-1 for "status" in fiber site

	app.Get("/fiber", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return nil
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).SendString("holly shit")
	})

	app.Get("/world", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("hello/hacker")
	})

	//testing-end

	type Movie struct { // struct name always start in captial letter
		Nameofmovie string `query:"name"`
	}

	app.Get("/parameter", func(c *fiber.Ctx) error {
		Movie := Movie{}
		if err := c.QueryParser(&Movie); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)

		}
		str := fmt.Sprintf("the name of the movie is : %s", Movie.Nameofmovie)
		return c.Send([]byte(str)) // every fuction and struct words defentely start with "capital letter" , but query dont't need this
	})

	type Series struct {
		Nameofseries string `json:"name"`
	}

	app.Post("/postdata", func(c *fiber.Ctx) error {
		Series := Series{}
		if err := c.JSON(&Series); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)

		}
		str := fmt.Sprintf("the name of the movie is : %s", Series.Nameofseries)
		return c.Send([]byte(str)) // every fuction and struct words defentely start with "capital letter" , but query dont't need this
	})

	app.Get("/name/:name/:another", func(c *fiber.Ctx) error {
		name := c.Params("name") //parameter
		another := c.Params("another")
		response := fmt.Sprintf("hello %s ! and hello %s !", name, another)
		return c.Send([]byte(response))
	})

	// fiber serverantworten - part-3

	app.Get("/forbbiden", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).Send([]byte("iam not a dump"))
	})

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	app.Get("/response/success", func(c *fiber.Ctx) error {
		return c.JSON(Response{true, "everything is a loop"})
	})

	app.Get("/response/fail", func(c *fiber.Ctx) error {
		return c.JSON(Response{false, "sometimes worst case also happen"})
	})

	// poral :

	app.Listen(":3000")

}
