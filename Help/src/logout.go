package main

import "net/http"

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// delete the cookie
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	// redirect to the home page
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
