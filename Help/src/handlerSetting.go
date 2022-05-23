package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** changePasswordHandler ***")
	d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	fmt.Println("oldPassword: ", oldPassword)
	fmt.Println("newPassword: ", newPassword)
	if CheckPasswordFromUserId(d.UserID, oldPassword) {
		if ChangePasswordFromUserId(w, d.UserID, newPassword) {
			http.Redirect(w, r, "/settingProfile", http.StatusFound)
		} else {
			d.AddPostError.Content = "Les critères de sécurité ne sont pas respectés"
			TMPL.ExecuteTemplate(w, "settingProfile.html", d)
		}
	} else {
		d.AddPostError.Content = "Les mots de passe ne correspondent pas"
		TMPL.ExecuteTemplate(w, "settingProfile.html", d)
	}
}

func CheckPasswordFromUserId(userId int, password string) bool {
	var passwordFromDB string
	err := DB.QueryRow("SELECT Password FROM users WHERE User_id = ?", userId).Scan(&passwordFromDB)
	if err != nil {
		panic(err.Error())
	}
	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDB), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func ChangePasswordFromUserId(w http.ResponseWriter, userId int, password string) bool {
	passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces := CheckPassword(password)
	if !passwordLowercase || !passwordUppercase || !passwordNumber || !passwordSpecial || !passwordLength || !passwordNoSpaces {
		return false
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	stmt, err := DB.Prepare("UPDATE users SET Password = ? WHERE User_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(hashedPassword, userId)
	if err != nil {
		panic(err.Error())
	}
	return true
}
