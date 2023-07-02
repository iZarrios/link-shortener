package db

import (
	"os"
	"strconv"

	"github.com/gofiber/storage/postgres/v2"
)

func Setup() *postgres.Storage {
	host := os.Getenv("DB_HOST")
	if host == "" {
		panic("DB_HOST is not set")
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		panic("DB_USER is not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		panic("DB_PASSWORD is not set")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		panic("DB_NAME is not set")
	}

	table := os.Getenv("DB_TABLE")
	if table == "" {
		panic("DB_TABLE is not set")
	}

	port_str := os.Getenv("DB_PORT")
	if port_str == "" {
		panic("DB_PORT is not set")
	}

	port, err := strconv.Atoi(port_str)
	if err != nil {
		panic("DB_PORT is not a number")
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		panic("DB_SSLMODE is not set")
	}

	cfg := postgres.Config{
		Host:     host,
		Port:     port,
		Username: user,
		Password: password,
		Database: dbname,
		SSLMode:  sslmode,
		Table:    table,
	}
	store := postgres.New(cfg)
	return store
}
