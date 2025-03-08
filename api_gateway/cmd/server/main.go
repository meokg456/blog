package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meokg456/api_gateway/internal/config"
	"github.com/meokg456/api_gateway/internal/middleware"
)

func setupRoutes(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {
	config.LoadConfig()

	router := mux.NewRouter()

	router.Use(middleware.AuthMiddleware)
	router.Use(middleware.ForwardMiddleware)

	setupRoutes(router)

	log.Println("Api gateway is listening on http://localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err.Error())
	}
}
