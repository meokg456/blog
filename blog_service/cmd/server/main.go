package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/meokg456/blog_service/internal/db"
	"github.com/meokg456/blog_service/internal/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, "\n", r.Body)
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes(router *mux.Router) {
	router.Use(loggingMiddleware)

	handler.SetupPostRoutes(router)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	db.InitDB()

	router := mux.NewRouter()
	SetupRoutes(router)

	log.Println("Blog service is listening on http://localhost:8082")

	if err := http.ListenAndServe("localhost:8082", router); err != nil {
		log.Fatal(err)
	}
}
