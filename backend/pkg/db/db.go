package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlStore struct {
	*gorm.DB
}

func NewSqlStore() *SqlStore {
	return &SqlStore{setup()}
}

func setup() *gorm.DB {
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
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)
	db, _ := gorm.Open(postgres.Open(dns), &gorm.Config{})
	return db

}
