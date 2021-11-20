package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Listening to port :8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/thread/read", readThread)
	mux.HandleFunc("/thread/post", postThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// data.CreateUser()

	server.ListenAndServe()
}
