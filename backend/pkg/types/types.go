package types

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Link struct {
	ID             uint64    `json:"id" gorm:"primary_key"`
	Url            string    `json:"url" gorm:"not null unique"`
	Redirect       string    `json:"redirect" gorm:"not null unique"`
	NumberOfVisits uint64    `json:"number_of_visits" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"not null default CURRENT_TIMESTAMP "`
}

type APISever struct {
	*fiber.App
	Port string
	Host string
}
