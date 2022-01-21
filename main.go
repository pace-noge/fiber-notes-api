package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pace-noge/fiber-notes-api/database"
	"github.com/pace-noge/fiber-notes-api/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
