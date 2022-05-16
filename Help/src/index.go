package main

import (
	"fmt"
	"net/http"
	"time"
)

type data struct {
	Name  string
	ID    []int
	SName []string
	Date  []*time.Time
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
	for i := 0; i < len(S); i++ {
		d.ID = append(d.ID, S[i].ID)
		d.SName = append(d.SName, S[i].Name)
		d.Date = append(d.Date, S[i].Date)
	}
	err = tmpl.ExecuteTemplate(w, "post.html", d)
	if err != nil {
		return
	}
}
