package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/SrijanSriv/gorum/data"
)

func login(w http.ResponseWriter, r *http.Request) {

	var templates *template.Template
	login_temp_files := []string{"templates/public.layout.html", "templates/public.navbar.html", "templates/login.html"}

	templates = template.Must(template.ParseFiles(login_temp_files...))

	templates.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	var templates *template.Template
	login_temp_files := []string{"templates/public.layout.html", "templates/public.navbar.html", "templates/signup.html"}

	templates = template.Must(template.ParseFiles(login_temp_files...))

	templates.Execute(w, nil)
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}
	user := data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if ok, _ := user.CheckUserExistance(); !ok {

		if err = user.CreateUser(); err != nil {
			log.Fatal()
		}
		http.Redirect(w, r, "/login", http.StatusFound)

	} else {
		fmt.Println("User already exists!") /*pop up*/
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		fmt.Println("error at email checking")
		log.Fatal(err)
	}

	if user.Password == r.PostFormValue("password") {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("error at session creation")
			log.Fatal(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		fmt.Println("No such user exists!") /*pop up*/
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}
