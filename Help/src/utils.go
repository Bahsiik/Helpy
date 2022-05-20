package main

import (
	"fmt"
	"net/http"
	"unicode"
)

func LoginError(w http.ResponseWriter, err error, errMsg string) bool {
	err = TMPL.ExecuteTemplate(w, "login.html", errMsg)
	if err != nil {
		fmt.Println("err: ", err)
		return true
	}
	return false
}

func GetUsernameFromSession(w http.ResponseWriter, r *http.Request) Data {
	cookie := CheckCookie(w, r)
	userID := SelectUserIDFromSessionID(cookie.Value)
	username := SelectUsernameFromID(userID)
	d := Data{
		Username: username,
	}
	return d
}

func GetUserIDFromSession(w http.ResponseWriter, r *http.Request) Data {
	cookie := CheckCookie(w, r)
	userID := SelectUserIDFromSessionID(cookie.Value)
	d := Data{
		UserID: userID,
	}
	return d
}

func CheckCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No cookie found")
		err := TMPL.ExecuteTemplate(w, "login.html", "Vous devez vous connecter pour accéder à cette page")
		if err != nil {
			fmt.Println("Error executing template")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
	}
	return cookie
}

func CheckUsername(username string) (bool, bool) {
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	return nameAlphaNumeric, nameLength
}

func CheckPassword(password string) (bool, bool, bool, bool, bool, bool) {
	var passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength bool
	passwordNoSpaces := true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			passwordLowercase = true
		case unicode.IsUpper(char):
			passwordUppercase = true
		case unicode.IsNumber(char):
			passwordNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			passwordSpecial = true
		case unicode.IsSpace(int32(char)):
			passwordNoSpaces = false
		}
	}
	if 11 < len(password) && len(password) < 60 {
		passwordLength = true
	}
	return passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces
}
