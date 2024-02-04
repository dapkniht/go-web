package controller

import (
	"net/http"
	"web-1/helper"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	helper.ClearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
}
