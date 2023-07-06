package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/types"
)

func SetupRedirectRoutes(app fiber.Router, db *db.SqlStore) {
	r := app.Group("/r")
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Redirect")
	})
	r.Get("/:url", Redirect(db))
}

func FindByUrl(url string, db *db.SqlStore) (*types.Link, error) {
	var link *types.Link
	tx := db.Where("redirect= ?", url).First(&link)
	return link, tx.Error
}

func Redirect(db *db.SqlStore) func(*fiber.Ctx) error {
	response := &ApiResponse{
		Code: fiber.StatusNotFound,
		Msg:  "Redirect not found",
		Data: nil,
	}
	return func(c *fiber.Ctx) error {
		url := c.Params("url")

		link, err := FindByUrl(url, db)
		if err != nil {
			return c.JSON(response)
		}
		link.NumberOfVisits += 1
		db.Save(&link)
		return c.Redirect("https://"+link.Url, fiber.StatusTemporaryRedirect)

	}
}
