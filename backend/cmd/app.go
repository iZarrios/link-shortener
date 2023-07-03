package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/routes"
)

type Server struct {
	*fiber.App
	Port string
	Host string
}

func NewServer(db *db.SqlStore) *Server {
    cfg := fiber.Config{
        DisableStartupMessage: false,
        Prefork:               false,
    }
	app := fiber.New(cfg)

	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(logger.New(logger.Config{
		Format:        "${pid} ${status} - ${method} ${path}\n",
		TimeFormat:    "02-Jan-2006",
		TimeZone:      "Egypt/Cairo",
		DisableColors: false,
	}))

	v1 := app.Group("/api/v1")

	routes.SetupLinksRoutes(v1, db)

	return &Server{
		App:  app,
		Port: os.Getenv("API_PORT"),
		Host: os.Getenv("API_HOST"),
	}
}
