package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// La fonction loginHandler est une fonction de gestionnaire qui est appelée lorsque l'utilisateur accède à la route
// /login. Il rend le modèle login.html
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Handler ****")
	err := tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Auth Handler ****")

	// On récupère les données du formulaire
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	// On récupère le hashage du mot de passe de l'utilisateur dans la base de données (à partir de son nom d'utilisateur)
	var hash string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hash)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du hashage du mot de passe de l'utilisateur")
		tmpl.ExecuteTemplate(w, "login.html", "Veuillez vérifier vos identifiants")
		return
	}

	// On vérifie que le mot de passe de l'utilisateur correspond au hashage en le hashant avec le même algorithme
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("Erreur lors de la vérification du mot de passe de l'utilisateur")
		tmpl.ExecuteTemplate(w, "login.html", "Veuillez vérifier vos identifiants")
		return
	} else {
		fmt.Fprintln(w, "Vous êtes connecté à ", username)
	}
}
