package controller

import (
	"html/template"
	"net/http"
)

func SettingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("views/setting.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		person := GetPerson()

		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		age := r.Form.Get("age")
		EditPerson(name, age)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}
