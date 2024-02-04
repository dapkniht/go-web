package controller

import (
	"context"
	"html/template"
	"net/http"
	"path/filepath"
	"web-1/db"
	"web-1/model"

	"go.mongodb.org/mongo-driver/bson"

	"web-1/helper"

	"go.mongodb.org/mongo-driver/mongo"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		currentDir, _ := filepath.Abs(filepath.Dir("."))
		temp, err := template.ParseFiles(filepath.Join(currentDir, "/views/auth/login.html"))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, nil)

	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")

		connect := db.DbConnect()

		var user model.User

		filter := bson.M{
			"email":    email,
			"password": password,
		}
		err := connect.Database("test").Collection("users").FindOne(context.Background(), filter).Decode(&user)
		if err == mongo.ErrNoDocuments {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			helper.SetSession(user.Email, user.Password, w, r)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}
