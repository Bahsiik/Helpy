package main

import (
	"fmt"
	"net/http"
)

type data struct {
	Name     string
	Subjects []Subject
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** indexHandler ****")
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No c found")
	}
	S := selectAllSubjects()
	userID := getUserIdFromSession(c.Value)
	username := getUsernameFromID(userID)
	d := data{
		Name:     username,
		Subjects: S,
	}
	err = tmpl.ExecuteTemplate(w, "post.html", d)
	if err != nil {
		return
	}
}
