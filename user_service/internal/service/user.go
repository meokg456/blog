package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/meokg456/user_service/internal/db"
	"github.com/meokg456/user_service/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(request *model.RegisterRequest) error {
	var existinguser string
	err := db.DB.QueryRow("SELECT username FROM users WHERE username = $1", request.Username).Scan(&existinguser)
	if err == nil {
		return errors.New("username is already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("server fault")
	}

	var userId int
	saveError := db.DB.
		QueryRow("INSERT INTO users (username, password, fullName) VALUES ($1, $2, $3) RETURNING id", request.Username, hashedPassword, request.FullName).
		Scan(&userId)

	if saveError != nil {
		return saveError
	}

	return nil
}

func Login(request *model.LoginRequest) (string, error) {
	var user model.User

	err := db.DB.QueryRowx("SELECT * FROM users WHERE username = $1", request.Username).StructScan(&user)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", errors.New("incorrect password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
