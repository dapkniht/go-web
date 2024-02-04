package middleware

import (
	"net/http"
	"web-1/helper"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helper.IsLoggedIn(r) {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func AuthLoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if helper.IsLoggedIn(r) {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
