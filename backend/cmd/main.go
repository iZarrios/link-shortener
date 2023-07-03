package main

import (
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/types"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db := db.NewSqlStore()
	if db == nil {
		panic("Error creating store")
	}

	db.AutoMigrate(&types.Link{})

	link := &types.Link{
		Url:      "https://www.google.com",
		Redirect: "no",
		Count:    0,
	}
	db.Save(&link)
}
