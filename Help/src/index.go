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
	// get the c named "session" using r.Cookie
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No c found")
	}
	fmt.Println("Cookie: ", c)
	fmt.Println("Cookie value: ", c.Value)
	fmt.Println("Cookie name: ", c.Name)
	fmt.Println("Cookie max age: ", c.MaxAge)
	// redirect to /test.html
	S := selectSubjects()
	// create a new data struct
	// mix the data of S and c.Value
	d := data{
		Name: c.Value,
	}
	// add the subjects to the data struct
	d.Subjects = S
	err = tmpl.ExecuteTemplate(w, "post.html", d)
	if err != nil {
		return
	}
}
