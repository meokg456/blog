package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

var DB *sqlx.DB

func InitDB() {
	db, err := sqlx.Open("postgres", "host=localhost user=postgres password=root dbname=blog_service sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
