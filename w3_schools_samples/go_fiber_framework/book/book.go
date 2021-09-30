package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/vigneshwaran/go_fiber_framework/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books) // c.Send("All books")
}
func GetBook(c *fiber.Ctx) {
	id := c.Params("id") //c.Send("A single book")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}
func NewBook(c *fiber.Ctx) {
	db := database.DBConn // c.Send("Add a new  books")
	var book Book
	book.Title = "1984"
	book.Author = "george orwell"
	book.Rating = 5

	db.Create(&book)
	c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id") // c.Send("delete a book")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("no book found with given ID")
		return
	}
	db.Delete(&book)
	c.Send("book sucessfully delete")

}
