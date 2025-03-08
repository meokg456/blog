package model

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	FullName string `db:"full_name"`
}
