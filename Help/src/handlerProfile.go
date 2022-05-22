package main

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileComHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileComHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileCom.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileFavHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileFavHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileFav.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileLikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileLikeHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileLike.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfilePubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profilePubHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profilePub.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
