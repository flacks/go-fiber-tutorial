package main

import (
	"fmt"
	"github.com/flacks/go-fiber-tutorial/book"
	"github.com/flacks/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully established")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database successfully migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
}

func main() {
	app := fiber.New()

	initDatabase()
	defer func() {
		if err := database.DBConn.Close(); err != nil {
			panic(err)
		}
	}()

	setupRoutes(app)

	if err := app.Listen(3000); err != nil {
		panic(err)
	}
}
