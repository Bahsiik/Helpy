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

func ChangeAvatarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** changeAvatarHandler ***")
	d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	d.Avatar = r.FormValue("avatar")
	ChangeAvatarFromUserId(d.UserID, d.Avatar)
	d.AvatarRoute = TranslateAvatarIdToString(d.Avatar)
	http.Redirect(w, r, "/settingProfile", http.StatusFound)
}

func ChangeAvatarFromUserId(userId int, avatar string) {
	stmt, err := DB.Prepare("UPDATE users SET Profil_Pic = ? WHERE User_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(avatar, userId)
	if err != nil {
		panic(err.Error())
	}
}

func TranslateAvatarIdToString(avatarId string) string {
	var avatar string
	switch avatarId {
	case "1":
		avatar = "profil.jpg"
	case "2":
		avatar = "profil_black.jpg"
	case "3":
		avatar = "profil_green.jpg"
	case "4":
		avatar = "profil_pink.jpg"
	case "5":
		avatar = "profil_red.jpg"
	case "6":
		avatar = "profil_yellow.jpg"
	}
	return avatar
}

func SettingNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingNotificationsHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingNotifications.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SettingAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingAccountHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingAccount.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
