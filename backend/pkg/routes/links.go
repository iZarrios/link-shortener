package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/types"
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

	var links []types.Link
	db.Find(&links)

	response := &ApiResponse{
		Code: fiber.StatusOK,
		Msg:  "Links retrieved successfully",
		Data: links,
	}
	return func(c *fiber.Ctx) error { // return of the function is a fiber handler
		c.Status(response.Code)
		return c.JSON(response) // return of the handler is an error

	}
}

func createLink(db *db.SqlStore) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("createLink")
	}
}

func getLink(db *db.SqlStore) func(*fiber.Ctx) error {
	response := &ApiResponse{
		Code: fiber.StatusOK,
		Msg:  "Link retrieved successfully",
		Data: nil,
	}
	return func(c *fiber.Ctx) error {
		id, err_conv := strconv.Atoi(c.Params("id"))

		var link *types.Link
		err := db.First(&link, id).Error
		if err != nil || err_conv != nil {
			response.Msg = "Link not found"
		} else {
			c.Status(response.Code)
			response.Data = link
		}
		return c.JSON(response)
	}
}

func updateLink(db *db.SqlStore) func(*fiber.Ctx) error {
    response := &ApiResponse{
        Code: fiber.StatusOK,
        Msg:  "Link updated successfully",
        Data: nil,
    }
	return func(c *fiber.Ctx) error {
        return c.JSON(response)
	}
}

func deleteLink(db *db.SqlStore) func(*fiber.Ctx) error {
    response := &ApiResponse{
        Code: fiber.StatusOK,
        Msg:  "Link deleted successfully",
        Data: nil,
    }
	return func(c *fiber.Ctx) error {
        id, err_conv := strconv.Atoi(c.Params("id"))
        err :=db.Delete(&types.Link{}, id).Error
        if err != nil || err_conv != nil {
            response.Msg = "Link not found"
        } else {
            c.Status(response.Code)
            response.Data = id 
        }
        return c.JSON(response)
	}
}
