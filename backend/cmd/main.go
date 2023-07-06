package main

import (
	"flag"
	"fmt"

	"github.com/iZarrios/link-shortener/backend/pkg/db"
	"github.com/iZarrios/link-shortener/backend/pkg/routes"
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

	toMigrate := flag.Bool("migrate", false, "set to true to migrate")
	flag.Parse()

	if *toMigrate {
		fmt.Println("Migrating...")

		// add new types here
		types := []interface{}{
			&types.Link{},
		}

		db.AutoMigrate(types...)
		fmt.Println("Migration complete")
	}

	app := NewServer(db)
	app.Use(routes.NotFoundRoute)

	app.Listen(fmt.Sprintf("%s:%s", app.Host, app.Port))
}
