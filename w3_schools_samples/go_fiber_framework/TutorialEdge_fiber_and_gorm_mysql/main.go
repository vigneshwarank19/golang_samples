package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vigneshwaran/fiber/gorm/mysql/book"
	"github.com/vigneshwaran/fiber/gorm/mysql/database"
)

func hellovk(c *fiber.Ctx) error {
	return c.Send([]byte("welocome to my vlog"))
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("database connection succesfully open")

	database.DBConn.AutoMigrate(&book.Book{})

	fmt.Println("Database Migrate")
}

func main() {
	app := fiber.New()

	initDatabase()

	defer database.DBConn.Close()

	app.Get("/", hellovk)

	SetupRoutes(app)

	app.Listen(":3000")
}
