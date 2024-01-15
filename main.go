package main

import (
	"log"
	"net/http"
	c "web-1/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/gallery", c.GalleryHandler).Methods("GET")
	r.HandleFunc("/gallery/{id}", c.ImageHandler).Methods("GET")
	r.HandleFunc("/about", c.AboutHandler).Methods("GET")
	r.HandleFunc("/setting", c.SettingHandler).Methods("GET")
	r.HandleFunc("/", c.IndexHandler).Methods("GET")

	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/assets"))))

	log.Println("Berjalan Pada Port 8000")
	http.ListenAndServe(":8000", nil)
}
