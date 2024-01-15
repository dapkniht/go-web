package controller

import (
	"html/template"
	"net/http"
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

	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}
