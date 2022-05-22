package main

import (
	"fmt"
	"net/http"
)

func ThematiquesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** thematiquesHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "thematiques.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** aboutHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "about.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** teamHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "team.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
