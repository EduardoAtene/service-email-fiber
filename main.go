package main

import (
	"database/sql"
	"fmt"

	"github.com/EduardoAtene/service-email-fiber/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	prefixEmailConst string = "email"
)

func main() {
	app := incializeApp()

	incializeDb()

	v1 := app.Group("v1")
	email := v1.Group(prefixEmailConst)
	email.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	email.Get("/:id?", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "Hello %s\n", c.Params("id"))

		return nil
	})

	app.Listen(":3000")

}

func incializeApp() (app *fiber.App) {
	app = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return
}

func incializeDb() *sql.DB {
	db, err := db.OpenConnection()
	if err != nil {
		panic(err)
	}

	return db
}
