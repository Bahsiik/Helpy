package main

import (
	"math/rand"
	"net/http"
	"time"
)

func generateSessionId() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AddSessionCookie(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sessionId := generateSessionId()
	userId := SelectUserIDFromUsername(username)
	date := time.Now()
	session := Session{
		Token:  sessionId,
		UserID: userId,
		Date:   &date,
	}
	saveSession(session)
	cookie := &http.Cookie{
		Name:    "session",
		Value:   sessionId,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}
	http.SetCookie(w, cookie)
}

func saveSession(session Session) {
	stmt := "INSERT INTO session (Session_id, User_id) VALUES (?, ?)"
	_, err := DB.Exec(stmt, session.Token, session.UserID)
	if err != nil {
		panic(err)
	}
}
