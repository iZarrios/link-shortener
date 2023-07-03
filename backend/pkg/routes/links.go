package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/link-shortener/backend/pkg/db"
)

func SetupLinksRoutes(app fiber.Router, db *db.SqlStore) {
	links := app.Group("/links")
	links.Get("/", getLinks(db))
	links.Post("/", createLink(db))
	links.Get("/:id", getLink(db))
	links.Put("/:id", updateLink(db))
	links.Delete("/:id", deleteLink(db))
}

func getLinks(db *db.SqlStore) func(c *fiber.Ctx) error {

	response := &ApiResponse{
		Code:  fiber.StatusOK,
		Msg:   "Links retrieved successfully",
		Error: true,
		Data:  nil,
	}
	return func(c *fiber.Ctx) error { // return of the function is a fiber handler
		return c.JSON(response)// return of the handler is an error

	}
}

func createLink(db *db.SqlStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("createLink")
	}
}

func getLink(db *db.SqlStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("getLink")
	}
}

func updateLink(db *db.SqlStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("updateLink")
	}
}

func deleteLink(db *db.SqlStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("deleteLink")
	}
}
