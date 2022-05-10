package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tmpl *template.Template

func main() {

	tmpl, _ = template.ParseGlob("templates/*.html")
	cssFolder := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "forum",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		println("1")
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		println("2")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}
