package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/meokg456/api_gateway/internal/config"
)

func ForwardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Println(config.Proxies)

		for prefix, proxy := range config.Proxies {
			if strings.HasPrefix(path, prefix) {
				r.Header.Add("X-Forwarded-For", r.RemoteAddr)
				r.Header.Add("X-API-Gateway", "true")
				proxy.ServeHTTP(w, r)
			}
		}

		next.ServeHTTP(w, r)
	})
}
