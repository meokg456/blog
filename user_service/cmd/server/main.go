package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/meokg456/user_service/internal/db"
	"github.com/meokg456/user_service/internal/handler"
)

func SetupRoutes(router *mux.Router) {
	handler.SetupUserRouter(router)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	db.InitDB()

	router := mux.NewRouter()

	SetupRoutes(router)

	log.Println("User service is listening on http://localhost:8081")

	if err := http.ListenAndServe("localhost:8081", router); err != nil {
		log.Fatal(err)
	}
}
