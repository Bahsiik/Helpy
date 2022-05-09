package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func addUser(user Users) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	result, err := db.Exec("INSERT INTO users (Username, Password, Passwordo, email) VALUES (?, ?, ?, ?)", user.Username, hashedPassword, user.Passwordo, user.Email)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	// Récupération de l'ID pour le return
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil

}

/*
func checkUser(user Users) (bool, error) {
	var username string
	var password string

	err := db.QueryRow("SELECT Username, Password FROM users WHERE Username = ?", user.Username).Scan(&username, &password)
	if err != nil {
		return false, fmt.Errorf("checkUser: %v", err)
	}
	if username == user.Username && password == user.Password {
		return true, nil
	}
	return false, nil
}
*/
