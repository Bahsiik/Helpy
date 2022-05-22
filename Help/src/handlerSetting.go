package main

import (
	"fmt"
	"net/http"
)

func SettingProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingProfileHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingProfile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
