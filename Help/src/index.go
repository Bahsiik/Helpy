package main

import (
	"fmt"
	"net/http"
)

type data struct {
	Name         string
	Posts        []Post
	AddPostError PostError
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** indexHandler ****")
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No c found")
	}
	S := selectAllPost()
	userID := getUserIdFromSession(c.Value)
	username := getUsernameFromID(userID)
	d := data{
		Name:  username,
		Posts: S,
	}
	err = tmpl.ExecuteTemplate(w, "index.html", d)
	if err != nil {
		return
	}
}
