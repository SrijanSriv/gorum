package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/SrijanSriv/gorum/data"
)

type Info struct {
	Threads  []data.Thread
	Username string
}

func index(w http.ResponseWriter, r *http.Request) {

	/*hardcoding input and parsing a slice*/
	// input := Info{"test", "Some Body", "srijan", time.Now().Format(time.ANSIC)}
	// inputAgain := Info{"testtesttesttesttest", "Some Body", "srijansriv", time.Now().Format(time.ANSIC)}
	// multiInput := []Info{input, inputAgain, input, inputAgain, input, inputAgain, input, inputAgain, input, inputAgain}
	/*above is all the hardcoded input*/
	threads := data.GetThreads()

	var templates *template.Template
	public_temp_files := []string{"templates/public.layout.html", "templates/index.html", "templates/public.navbar.html", "templates/aside.html"}
	private_temp_files := []string{"templates/private.layout.html", "templates/index.html", "templates/private.navbar.html", "templates/aside.html"}

	sess, err := session(w, r)
	_ = Info{threads, sess.Name}
	if err != nil {
		templates = template.Must(template.ParseFiles(public_temp_files...))
	} else {
		templates = template.Must(template.ParseFiles(private_temp_files...))
	}

	templates.Execute(w, threads)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			log.Fatal()
		}
	}
	return
}

func logout(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("_cookie")
	if err == nil {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}

	http.Redirect(w, r, "/", http.StatusFound)

}
