package noteRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "github.com/pace-noge/fiber-notes-api/internals/handlers/note"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/notes")

	note.Post("/", noteHandler.CreateNote)

	note.Get("/", noteHandler.GetNotes)

	note.Get("/:noteID", noteHandler.GetNote)

	note.Put("/:noteID", noteHandler.UpdateNote)

	note.Delete("/:noteID", noteHandler.DeleteNote)
}
