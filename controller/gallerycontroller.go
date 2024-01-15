package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type images struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

var Images []images

func getImages(w http.ResponseWriter) {
	apiURL := "https://api.thecatapi.com/v1/images/search?limit=50"
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := json.Unmarshal(body, &Images); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	if len(Images) == 0 {
		getImages(w)
	}

	if r.URL.Query().Get("refresh") == "true" {
		getImages(w)
		http.Redirect(w, r, "/gallery", http.StatusSeeOther)
	}

	tmp, err := template.ParseFiles("views/gallery.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmp.ExecuteTemplate(w, "gallery", Images)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageID := vars["id"]
	var image map[string]string
	image = map[string]string{}

	for _, value := range Images {
		if value.Id == imageID {
			image["id"] = value.Id
			image["url"] = value.Url
			break
		}
	}
	tmp, err := template.ParseFiles("views/gallery.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmp.ExecuteTemplate(w, "image", image)

}
