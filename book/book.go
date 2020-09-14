package book

import (
	"github.com/flacks/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book

	db.Find(&books)
	if len(books) == 0 {
		c.Send("No books found")
		return
	}

	if err := c.JSON(&books); err != nil {
		c.Status(500).Send("Failed to generate JSON")
		return
	}
}

func GetBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book

	db.Find(&book, c.Params("id"))
	if book.Title == "" {
		c.Status(400).Send("Book not found")
		return
	}

	if err := c.JSON(&book); err != nil {
		c.Status(500).Send("Failed to generate JSON")
		return
	}
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(400).Send("Error parsing book: ", err)
		return
	}

	db.Create(&book)

	if err := c.JSON(&book); err != nil {
		c.Status(500).Send("Failed to generate JSON")
		return
	}
}

func DeleteBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book

	db.First(&book, c.Params("id"))
	if book.Title == "" {
		c.Status(400).Send("Book not found")
		return
	}

	db.Delete(&book)

	if err := c.JSON(&book); err != nil {
		c.Status(500).Send("Failed to generate JSON")
		return
	}
}

func UpdateBook(c *fiber.Ctx) {
	db := database.DBConn
	var book Book

	db.Find(&book, c.Params("id"))
	if book.Title == "" {
		c.Status(400).Send("Book not found")
		return
	}

	if err := c.BodyParser(&book); err != nil {
		c.Status(400).Send("Error parsing book: ", err)
		return
	}

	db.Save(&book)

	if err := c.JSON(&book); err != nil {
		c.Status(500).Send("Failed to generate JSON")
		return
	}
}
