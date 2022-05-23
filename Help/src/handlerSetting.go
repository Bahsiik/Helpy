package main

import (
	"fmt"
	"net/http"
)

func SettingProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingProfileHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingProfile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ChangeUsernameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** changeUsernameHandler ***")
	d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	username := r.FormValue("newUsername")
	fmt.Println("username: ", username)
	fmt.Println("d: ", d.UserID)
	ChangeUsernameFromUserId(d.UserID, username)
	http.Redirect(w, r, "/settingProfile", http.StatusFound)
}

func ChangeUsernameFromUserId(userId int, username string) {
	stmt, err := DB.Prepare("UPDATE users SET Username = ? WHERE User_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(username, userId)
	if err != nil {
		panic(err.Error())
	}
}
