package main

import (
	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
        panic("Error loading .env file")
	}
	store := db.Setup()
	defer store.Close()
}
