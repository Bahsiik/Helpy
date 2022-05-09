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

	http.HandleFunc("/", FirstCookies)
	//CompareForm(w, r)
	//CompareCheck(w, r)
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}

func FirstCookies(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Test-cookie")
	fmt.Println("Cookie:", cookie, "Error:", err)
	if err != nil {
		fmt.Println("Cookies not found")
		cookie = &http.Cookie{
			Name:     "Test-cookie",
			Value:    "Test-value",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	erro := tmpl.ExecuteTemplate(w, "signup.html", nil)
	if erro != nil {
		return
	}
}

func CompareCheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	var user Users
	var err error
	err = db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		fmt.Fprintln(w, "Username or password is incorrect")
		return
	}
	if user.Password != password {
		fmt.Fprintln(w, "Username or password is incorrect")
		return
	}
	fmt.Fprintln(w, "You are logged in")
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func CompareForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	if r.FormValue("password") != r.FormValue("passwordo") {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		albID, err := addUser(Users{
			Username:  r.FormValue("username"),
			Password:  r.FormValue("password"),
			Passwordo: r.FormValue("passwordo"),
			Email:     r.FormValue("email"),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added Users: %v\n", albID)
	}
}
