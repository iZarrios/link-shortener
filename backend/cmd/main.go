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
		/////////////////////////
		types := []interface{}{
			&types.Link{},
		}
		/////////////////////////
		db.AutoMigrate(types...)
		fmt.Println("Migration complete")
	}

	app := NewServer(db)
    // setting up middlewares


    // setting up routes
	// swagger
	// public
	// private
	// not found
	app.Use(routes.NotFoundRoute)

	app.Listen(fmt.Sprintf("%s:%s", app.Host, app.Port))
	// rows, err := db.DB.Raw("SELECT * FROM links").Rows()
	// if err != nil {
	//     panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	//     var id uint64
	//     var url string
	//     var redirect string
	//     var count uint64
	//     var created_at string
	//     var updated_at string
	//     err = rows.Scan(&id, &url, &redirect, &count, &created_at, &updated_at)
	//     if err != nil {
	//         panic(err)
	//     }
	//     println(id, url, redirect, count, created_at, updated_at)
	// }

	// get all records in links table
	// var links []types.Link
	// db.Find(&links)
	// for _, link := range links {
	//
	//     // created_at := time.Time(link.CreatedAt).Format("2006-01-02 15:04:05")
	//     // updated_at := time.Time(link.UpdatedAt).Format("2006-01-02 15:04:05")
	//     link.CreatedAt = time.Time(link.CreatedAt)
	//     link.UpdatedAt = time.Time(link.UpdatedAt).Add(time.Hour * 8)
	//     fmt.Print(link.CreatedAt, link.UpdatedAt)
	//
	//     // println(link.ID, link.Url, link.Redirect, link.Count, created_at, updated_at)
	// }
}
