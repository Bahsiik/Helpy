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
	cookie := checkCookie(w, r)
	userID := getUserIdFromSession(cookie.Value)
	username := getUsernameFromID(userID)
	S := selectAllPost()
	d := data{
		Name:  username,
		Posts: S,
	}
	err := tmpl.ExecuteTemplate(w, "index.html", d)
	if err != nil {
		return
	}
}

func checkCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No cookie found")
		// redirect to login
		err := tmpl.ExecuteTemplate(w, "login.html", "Vous devez vous connec")
		if err != nil {
			fmt.Println("Error executing template")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
	}
	return cookie
}
