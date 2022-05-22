package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var TMPL *template.Template

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterHandler ****")
	err := TMPL.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}

func RegisterAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterAuthHandler ****")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	username := r.FormValue("username")
	firstPassword := r.FormValue("firstPassword")
	secondPassword := r.FormValue("secondPassword")
	nameAlphaNumeric, nameLength := CheckUsername(username)
	if firstPassword != secondPassword {
		fmt.Println("**** Passwords don't match ****")
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Les mots de passe ne correspondent pas")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces := CheckPassword(firstPassword)
	if !passwordLowercase || !passwordUppercase || !passwordNumber || !passwordSpecial || !passwordLength || !passwordNoSpaces || !nameAlphaNumeric || !nameLength {
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Veuillez vérifier les critères de nom d'utilisateur ou de votre mot de passe")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	stmt := "SELECT Username FROM users WHERE Username = ?"
	row := DB.QueryRow(stmt, username)
	var uID string
	err = row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Ce nom d'utilisateur existe déjà")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(firstPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	var insertStmt *sql.Stmt
	insertStmt, err = DB.Prepare("INSERT INTO users (Username, Password) VALUES (?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(username, hash)
	if err != nil {
		fmt.Println("error inserting new user:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Handler ****")
	err := TMPL.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Auth Handler ****")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	var hash string
	err = DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hash)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du hash du mot de passe de l'utilisateur: ", err)
		done := LoginError(w, err, "Veuillez vérifier vos identifiants")
		if done {
			return
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		done := LoginError(w, err, "Veuillez vérifier vos identifiants")
		if done {
			return
		}
		return
	} else {
		AddSessionCookie(w, r)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** indexHandler ****")
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** homeHandler ****")
	d := GetUserInfoFromSession(w, r)
	S := SelectAllPost()
	d.Posts = S
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
	}
	d.Topic = "Toutes les publications"
	d.TopicShortName = "all"
	err := TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}
