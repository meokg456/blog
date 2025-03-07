package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	db, err := sqlx.Open("postgres", "host=localhost user=postgres password=root dbname=blog_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func CloseDB() {
	DB.Close()
}
