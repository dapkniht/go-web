package helper

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func SetSession(username, password string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	session.Values["username"] = username
	session.Values["password"] = password
	session.Options = &sessions.Options{
		MaxAge: 60,
		Path:   "/",
	}
	session.Save(r, w)
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login-session")
	session.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}
	session.Save(r, w)
}

func IsLoggedIn(r *http.Request) bool {
	session, _ := store.Get(r, "login-session")
	username := session.Values["username"]
	password := session.Values["password"]

	return username != nil && password != nil
}
