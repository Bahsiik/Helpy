package main

import (
	"fmt"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** aboutHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "about.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** teamHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "team.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func settingProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingProfileHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "settingProfile.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func settingNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingNotificationsHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "settingNotifications.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func settingAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingAccountHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "settingAccount.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileHandler ***")
	cookie := checkCookie(w, r)
	err := tmpl.ExecuteTemplate(w, "profile.html", cookie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
