package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Session struct {
	Token  string
	UserID int
	Date   *time.Time
}

func generateSessionId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func addSessionCookie(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sessionId := generateSessionId()
	userId := getUserId(username)
	date := time.Now()
	// create a new session
	session := Session{
		Token:  sessionId,
		UserID: userId,
		Date:   &date,
	}
	// save the session in the database
	saveSession(session)
	fmt.Println("User id: ", userId)
	cookie := &http.Cookie{
		Name:    "session",
		Value:   sessionId,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}
	// On ajoute le cookie à la réponse
	http.SetCookie(w, cookie)
	// On vérifie que le cookie est bien présent dans la requête
	http.Redirect(w, r, "/index", http.StatusFound)
}

// Get the user id from the database from the username  (if the user exists)
func getUserId(username string) int {
	stmt := "SELECT User_id FROM users WHERE Username = ?"
	rows, err := db.Query(stmt, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

func saveSession(session Session) {
	stmt := "INSERT INTO session (Session_id, User_id) VALUES (?, ?)"
	_, err := db.Exec(stmt, session.Token, session.UserID)
	if err != nil {
		panic(err)
	}
}

func getUserIdFromSession(sessionID string) int {
	stmt := "SELECT User_id FROM session WHERE Session_id = ?"
	rows, err := db.Query(stmt, sessionID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

func getUsernameFromID(userID int) string {
	stmt := "SELECT Username FROM users WHERE User_id = ?"
	rows, err := db.Query(stmt, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
	}
	return username
}
