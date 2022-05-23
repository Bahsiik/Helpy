package main

import (
	"fmt"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** AdminHandler ****")
	d := GetUserInfoFromSession(w, r)
	if d.IsAdmin == false {
		http.Redirect(w, r, "/", http.StatusFound)
	}
	Users := SelectAllUsers()
	d.Users = Users
	err := TMPL.ExecuteTemplate(w, "moderation.html", d)
	if err != nil {
		return
	}
}

func MuteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** MuteUserHandler ***")
	r.ParseForm()
	userName := r.FormValue("username")
	SetUserMute(userName)
	http.Redirect(w, r, "/admin", http.StatusFound)
}

func SetUserMute(userName string) {
	stmt, err := DB.Prepare("UPDATE users SET Muted = 1 WHERE Username = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(userName)
	if err != nil {
		panic(err.Error())
	}
}

func UnmuteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** UnmuteUserHandler ***")
	r.ParseForm()
	userName := r.FormValue("username")
	SetUserUnmute(userName)
	http.Redirect(w, r, "/admin", http.StatusFound)
}

func SetUserUnmute(userName string) {
	stmt, err := DB.Prepare("UPDATE users SET Muted = 0 WHERE Username = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(userName)
	if err != nil {
		panic(err.Error())
	}
}
