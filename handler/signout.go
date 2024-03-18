package handler

import (
	"net/http"
	"time"
)

func Signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "auth_token",
		Expires: time.Now(),
	})
}
