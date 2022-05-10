package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// La fonction registerHandler est appelée lorsque l'utilisateur clique sur le lien d'inscription dans la barre de
// navigation. La fonction exécute ensuite le modèle register.html
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterHandler ****")
	err := tmpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		return
	}
}

// Il vérifie si le nom d'utilisateur et le mot de passe répondent aux critères, puis il vérifie si le nom d'utilisateur
// existe déjà dans la base de données, puis il crée un hachage du mot de passe, puis il insère le nom d'utilisateur et le
// hachage dans la base de données
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterAuthHandler ****")
	// récupération des données du formulaire
	err := r.ParseForm()
	if err != nil {
		return
	}

	username := r.FormValue("username")
	firstPassword := r.FormValue("firstPassword")
	secondPassword := r.FormValue("secondPassword")

	// vérification des critères de nom d'utilisateur
	nameAlphaNumeric, nameLength := checkUsername(username)

	// vérifier si les mots de passe sont identiques
	if firstPassword != secondPassword {
		fmt.Println("**** Passwords don't match ****")
		err = tmpl.ExecuteTemplate(w, "register.html", "⚠ Les mots de passe ne correspondent pas")
		if err != nil {
			return
		}
		return
	}

	// vérification des critères de mot de passe
	passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces := checkPassword(firstPassword)
	if !passwordLowercase || !passwordUppercase || !passwordNumber || !passwordSpecial || !passwordLength || !passwordNoSpaces || !nameAlphaNumeric || !nameLength {
		err = tmpl.ExecuteTemplate(w, "register.html", "⚠ Veuillez vérifier les critères de nom d'utilisateur ou de votre mot de passe")
		if err != nil {
			return
		}
		return
	}

	// vérifier si le nom d'utilisateur existe déjà dans la base de données
	stmt := "SELECT Username FROM users WHERE Username = ?"
	row := db.QueryRow(stmt, username)
	var uID string
	err = row.Scan(&uID)
	// si le nom d'utilisateur existe déjà, on affiche un message d'erreur
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		err = tmpl.ExecuteTemplate(w, "register.html", "⚠ Ce nom d'utilisateur existe déjà")
		if err != nil {
			return
		}
		return
	}

	// créer un hash de mot de passe
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(firstPassword), bcrypt.DefaultCost)
	// si le hash n'a pas pu être créé, on affiche un message d'erreur
	if err != nil {
		fmt.Println("bcrypt err:", err)
		err := tmpl.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			return
		}
		return
	}
	// DEBUG fmt.Println("hash:", hash)
	// DEBUG fmt.Println("string(hash):", string(hash))

	// insérer le nom d'utilisateur et le hash dans la base de données
	var insertStmt *sql.Stmt
	insertStmt, err = db.Prepare("INSERT INTO users (Username, Password) VALUES (?, ?);")
	// si l'insertion n'a pas pu être effectuée, on affiche un message d'erreur
	if err != nil {
		fmt.Println("error preparing statement:", err)
		err := tmpl.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			return
		}
		return
	}

	defer insertStmt.Close()
	// DEBUG (si utilisé, remplacer le "_" en dessous par "result") var result sql.Result
	// effectuer l'insertion
	_, err = insertStmt.Exec(username, hash)
	/* DEBUG
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	*/
	// si l'insertion n'a pas pu être effectuée, on affiche un message d'erreur
	if err != nil {
		fmt.Println("error inserting new user")
		err := tmpl.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			return
		}
		return
	}
	// si tout s'est bien passé, on affiche un message de succès
	fmt.Fprint(w, "congrats, your account has been successfully created")
}

// Il renvoie vrai si le nom d'utilisateur est alphanumérique et faux s'il ne l'est pas
func checkUsername(username string) (bool, bool) {
	// vérifier si le nom d'utilisateur est alphanumérique
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			// si le nom d'utilisateur n'est pas alphanumérique, on renvoie faux
			nameAlphaNumeric = false
		}
	}
	// vérifier si le nom d'utilisateur est assez long
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		// si le nom d'utilisateur est assez long, on renvoie vrai (de base nameLength = false)
		nameLength = true
	}
	return nameAlphaNumeric, nameLength
}

// Il vérifie si une chaîne contient au moins une lettre minuscule, une lettre majuscule, un chiffre, un caractère spécial,
// fait entre 11 et 60 caractères et ne contient pas d'espaces
func checkPassword(password string) (bool, bool, bool, bool, bool, bool) {
	// DEBUG fmt.Println("password:", password, "\npswdLength:", len(password))
	// vérifier si le mot de passe contient au moins une lettre minuscule, une lettre majuscule, un chiffre et un caractère spécial
	var passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength bool
	passwordNoSpaces := true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			// si le mot de passe contient une lettre minuscule, on renvoie vrai
			passwordLowercase = true
		case unicode.IsUpper(char):
			// si le mot de passe contient une lettre majuscule, on renvoie vrai
			passwordUppercase = true
		case unicode.IsNumber(char):
			// si le mot de passe contient un chiffre, on renvoie vrai
			passwordNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			// si le mot de passe contient un caractère spécial, on renvoie vrai
			passwordSpecial = true
		case unicode.IsSpace(int32(char)):
			// si le mot de passe contient un espace, on renvoie faux
			passwordNoSpaces = false
		}
	}
	// vérifier si le mot de passe est assez long
	if 11 < len(password) && len(password) < 60 {
		// si le mot de passe est assez long, on renvoie vrai (de base passwordLength = false)
		passwordLength = true
	}
	// DEBUG fmt.Println("passwordLowercase:", passwordLowercase, "\npasswordUppercase:", passwordUppercase, "\npasswordNumber:", passwordNumber, "\npasswordSpecial:", passwordSpecial, "\npasswordLength:", passwordLength, "\npasswordNoSpaces:", passwordNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	return passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces
}
