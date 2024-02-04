package main

import (
	"log"
	"net/http"
	c "web-1/controller"

	ac "web-1/controller/auth"

	m "web-1/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	r.Use(m.LogMiddleware)

	r.Handle("/login", m.AuthLoginMiddleware(http.HandlerFunc(ac.LoginHandler))).Methods("GET", "POST")
	r.Handle("/logout", m.AuthMiddleware(http.HandlerFunc(ac.LogoutHandler))).Methods("GET")
	r.Handle("/gallery", m.AuthMiddleware(http.HandlerFunc(c.GalleryHandler))).Methods("GET")
	r.Handle("/gallery/{id}", m.AuthMiddleware(http.HandlerFunc(c.GalleryHandler))).Methods("GET")
	r.Handle("/about", m.AuthMiddleware(http.HandlerFunc(c.AboutHandler))).Methods("GET")
	r.Handle("/setting", m.AuthMiddleware(http.HandlerFunc(c.SettingHandler))).Methods("GET", "POST")
	r.Handle("/", m.AuthMiddleware(http.HandlerFunc(c.IndexHandler))).Methods("GET")

	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/assets"))))

	log.Println("Berjalan Pada Port 8000")
	http.ListenAndServe(":8000", nil)
}
