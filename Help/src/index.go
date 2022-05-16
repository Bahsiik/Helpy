package main

import (
	"fmt"
	"net/http"
)

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
	err = tmpl.ExecuteTemplate(w, "post.html", c.Value)
	if err != nil {
		return
	}
}
