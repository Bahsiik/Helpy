package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func addUser(use User) (int64, error) {
	// Création de la requête SQL

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(use.Password), 8)
	result, err := db.Exec("INSERT INTO users (Username, Password, Passwordo, email) VALUES (?, ?, ?, ?)", use.Username, hashedPassword, use.Passwordo, use.Email)
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
