package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// La fonction loginHandler est une fonction de gestionnaire qui est appelée lorsque l'utilisateur accède à la route
// /login. Il rend le modèle login.html
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Handler ****")
	err := tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Nous obtenons le nom d'utilisateur et le mot de passe du formulaire, nous obtenons le hachage du mot de passe de la base
// de données, nous comparons le hachage du mot de passe avec le hachage du mot de passe que nous avons obtenu de la base
// de données, et s'ils correspondent, nous affichons un message disant que l'utilisateur est connecté
func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Auth Handler ****")
	// On récupère les données du formulaire
	err := r.ParseForm()
	if err != nil {
		// DEBUG fmt.Println("err: ", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	// On récupère le hashage du mot de passe de l'utilisateur dans la base de données (à partir de son nom d'utilisateur)
	var hash string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hash)
	// Si l'utilisateur n'existe pas, on affiche un message d'erreur
	if err != nil {
		// DEBUG fmt.Println("Erreur lors de la récupération du hashage du mot de passe de l'utilisateur: ", err)
		err = tmpl.ExecuteTemplate(w, "login.html", "Veuillez vérifier vos identifiants")
		if err != nil {
			// DEBUG fmt.Println("err: ", err)
			return
		}
		return
	}
	// On vérifie que le mot de passe de l'utilisateur correspond au hashage en le hashant avec le même algorithme
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Si le mot de passe ne correspond pas, on affiche un message d'erreur
	if err != nil {
		// DEBUG fmt.Println("Erreur lors de la vérification du mot de passe de l'utilisateur: ", err)
		err = tmpl.ExecuteTemplate(w, "login.html", "Veuillez vérifier vos identifiants")
		if err != nil {
			// DEBUG fmt.Println("err: ", err)
			return
		}
		return
	} else {
		addSessionCookie(w, r, username)
	}
}

func addSessionCookie(w http.ResponseWriter, r *http.Request, username string) {
	cookie := &http.Cookie{
		Name:    "session",
		Value:   username,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}
	// On ajoute le cookie à la réponse
	http.SetCookie(w, cookie)
	// On vérifie que le cookie est bien présent dans la requête
	http.Redirect(w, r, "/index", http.StatusFound)
}
