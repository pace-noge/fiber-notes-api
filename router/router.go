package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "github.com/pace-noge/fiber-notes-api/internals/routes/note"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)
}
