package main

import (
	"log"
	"net/http"

	"github.com/11010111/goauth/handler"
)

func main() {
	http.HandleFunc("/signin", handler.Signin)
	http.HandleFunc("/welcome", handler.Welcome)
	http.HandleFunc("/refresh", handler.Refresh)
	http.HandleFunc("/signout", handler.Signout)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
