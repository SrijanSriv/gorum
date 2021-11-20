package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening to port :", port)

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}

	// data.CreateUser()

	server.ListenAndServe()
}
