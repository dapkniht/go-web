package controller

import (
	"html/template"
	"net/http"
	"time"
)

var person = map[string]string{
	"name": "joni",
	"age":  "13",
}

func GetPerson() map[string]string {
	return person
}

func EditPerson(name, age string) {
	person["name"] = name
	person["age"] = age
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:    "examplecookie",
		Value:   "cookievalue",
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
	}

	// Menetapkan cookie di respons HTTP
	http.SetCookie(w, &cookie)

	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
