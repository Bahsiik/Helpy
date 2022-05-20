package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

func main() {
	TMPL, _ = template.ParseGlob("templates/*.html")
	cssFolder := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

	imgFolder := http.FileServer(http.Dir("img"))
	http.Handle("/img/", http.StripPrefix("/img/", imgFolder))

	//handle js
	jsFolder := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", jsFolder))
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "forum",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/registerauth", RegisterAuthHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/loginauth", LoginAuthHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/subjectByTopic", SelectPostTopicHandler)
	http.HandleFunc("/post", PostHandler)
	http.HandleFunc("/addPost", AddPostHandler)

	http.HandleFunc("/reply", ReplyHandler)
	http.HandleFunc("/addReply", AddReplyToPostHandler)
	http.HandleFunc("/team", TeamHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/profile", ProfileHandler)
	http.HandleFunc("/settingProfile", SettingProfileHandler)
	http.HandleFunc("/settingAccount", SettingAccountHandler)
	http.HandleFunc("/settingNotifications", SettingNotificationsHandler)
	http.HandleFunc("/postFeed", PostFeedHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		// DEBUG fmt.Println("err: ", err)
		return
	}
}
