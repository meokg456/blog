package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var UserProxy *httputil.ReverseProxy
var BlogProxy *httputil.ReverseProxy

func InitProxy() {
	userServiceUrl, _ := url.Parse("http://localhost:8081")
	UserProxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = userServiceUrl.Scheme
			req.URL.Host = userServiceUrl.Host
			req.Host = userServiceUrl.Host
		},
	}

	blogServiceUrl, _ := url.Parse("http://localhost:8082")
	BlogProxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = blogServiceUrl.Scheme
			req.URL.Host = blogServiceUrl.Host
			req.Host = blogServiceUrl.Host
		},
	}
}
