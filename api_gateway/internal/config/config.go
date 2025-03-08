package config

import (
	"log"
	"net/http/httputil"
	"os"

	"github.com/joho/godotenv"
	"github.com/meokg456/api_gateway/internal/proxy"
)

var JwtSecret string

var Proxies = map[string]*httputil.ReverseProxy{}

var PublicRoutes = map[string]bool{
	"/login":    true,
	"/register": true,
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	JwtSecret = os.Getenv("JWT_SECRET")

	proxy.InitProxy()

	Proxies["/login"] = proxy.UserProxy
	Proxies["/register"] = proxy.UserProxy
	Proxies["/post"] = proxy.BlogProxy
}
