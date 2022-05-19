package main

import (
	"fmt"
	"net/http"
)

type data struct {
	Username     string
	FirstPost    Post
	Posts        []Post
	AddPostError PostError
	Replies      []Reply
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** homeHandler ****")
	cookie := checkCookie(w, r)
	userID := getUserIdFromSession(cookie.Value)
	username := getUsernameFromID(userID)
	S := selectAllPost()
	d := data{
		Username: username,
		Posts:    S,
	}
	err := tmpl.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** indexHandler ****")
	// redirect to home
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func checkCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No cookie found")
		// redirect to login
		err := tmpl.ExecuteTemplate(w, "login.html", "Vous devez vous connecter pour accéder à cette page")
		if err != nil {
			fmt.Println("Error executing template")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
	}
	return cookie
}
