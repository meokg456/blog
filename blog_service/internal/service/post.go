package service

import (
	"errors"

	"github.com/meokg456/blog_service/internal/db"
)

func GetPosts() ([]db.Post, error) {
	var posts []db.Post
	err := db.DB.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func CreatePost(title string, content string) (*db.Post, error) {
	var post db.Post
	err := db.DB.QueryRowx("INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING *", title, content).StructScan(&post)

	if err != nil {
		return nil, err
	}

	return &post, err
}

func GetPost(id int) (*db.Post, error) {
	var post db.Post
	err := db.DB.Get(&post, "SELECT * FROM posts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func DeletePost(id int) error {
	result, err := db.DB.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Post is not existed")
	}

	return nil
}
