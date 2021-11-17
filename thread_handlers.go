package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/SrijanSriv/gorum/data"
)

type ThreadsPosts struct {
	Thread data.Thread
	Post   []data.Post
}

func readThread(w http.ResponseWriter, r *http.Request) {

	uuid := r.URL.Query().Get("id")
	stuff := ThreadsPosts{}
	post := []data.Post{{
		Id:        0,
		Uuid:      "0",
		Body:      "some body",
		UserId:    0,
		ThreadId:  0,
		CreatedAt: time.Now().Format("02-01-2006 15:04:05"),
	}}
	thread, err := data.ThreadByUuid(uuid)

	stuff.Thread = thread
	stuff.Post = post
	if err != nil {
		log.Fatal(err)
	} else {
		_, err := session(w, r)
		var templates *template.Template
		public_thread_files := []string{"templates/public.layout.html", "templates/public.navbar.html", "templates/public.thread.html"}
		private_thread_files := []string{"templates/private.layout.html", "templates/private.navbar.html", "templates/private.thread.html"}
		if err != nil {
			// fmt.Println(stuff)
			templates = template.Must(template.ParseFiles(public_thread_files...))
		} else {
			templates = template.Must(template.ParseFiles(private_thread_files...))
		}
		templates.Execute(w, stuff)
	}
}

func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		var templates *template.Template
		thread_files := []string{"templates/private.layout.html", "templates/new.thread.html", "templates/private.navbar.html"}

		templates = template.Must(template.ParseFiles(thread_files...))

		templates.Execute(w, nil) /*add structs instead of nil*/

	}
}

func createThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		user, err := data.UserByUuid(sess.Uuid)
		if err != nil {
			log.Fatal(err)
		}
		topic := r.PostFormValue("topic")
		body := r.PostFormValue("body")
		if err := data.CreateThread(topic, body, user.Name); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func postThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		user, err := data.UserByUuid(sess.Uuid)
		if err != nil {

			log.Fatal(err)
		}

		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := data.ThreadByUuid(uuid)

		if err != nil {
			log.Fatal(err)
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			log.Fatal(err)
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, r, url, http.StatusFound)
	}
}
