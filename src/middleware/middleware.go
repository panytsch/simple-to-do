package middleware

import (
	"net/http"
	"simple-to-do/src/helpers"
	"strings"
)

//AllowedPrefixes - list of allowed by middleware route prefixes
var AllowedPrefixes = []string{
	"api",
	"ws",
}

//IsThisRouteAllowedMiddleware - middleware to check is this route allowed
func IsThisRouteAllowedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if IsIfAllowedRequest(r.RequestURI) == false {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

func isItFileRequest(url string) bool {
	st := strings.Split(url, "/")
	fl := strings.Split(st[len(st)-1], ".")
	if len(fl) > 1 && len(fl[len(fl)-1]) > 0 { // example fl = ["test", "txt"] length 3
		return true
	}
	return false
}

//IsIfAllowedRequest checking url to awoid conflict of routers go and react
func IsIfAllowedRequest(url string) bool {
	if url == "/" {
		return true
	}
	return isRouteInAllowedPrefixes(url) || isItFileRequest(url)
}

func isRouteInAllowedPrefixes(url string) bool {
	r := strings.Split(url, "/")[1]
	return helpers.ContainsString(AllowedPrefixes, r)
}
