package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/types"
)

func FindByUrl(url string, db *db.SqlStore) (*types.Link, error) {
	var link *types.Link
	tx := db.Where("redirect= ?", url).First(&link)
	return link, tx.Error
}


func Redirect(db *db.SqlStore) func(*fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        url := c.Params("url")
        link, err := FindByUrl(url, db)
        if err != nil {
            return err
        }
        link.Count += 1
        db.Save(&link)
        return c.Redirect(link.Redirect)
    }
}
