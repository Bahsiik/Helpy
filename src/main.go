package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

// Ça je ne sais pas trop, mais c'est ce qui nous permet de manip la bdd
var db *sql.DB

// Ça je ne sais pas trop, mais c'est ce qui nous permet de manip la bdd

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	Password  string `json:"password"`
	Passwordo string `json:"passwordo"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
	Status    bool   `json:"status"`
	ProfilPic string `json:"profil_pic"`
}

func main() {

	// Gestion de tous les fichiers gohtml
	tmpl := template.Must(template.ParseGlob("../templates/login.gohtml"))
	cssFolder := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))

	// Paramètres de connexion à la BDD
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "forum",
		AllowNativePasswords: true,
	}
	// Création du handler de la BDD
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		println("1")
		log.Fatal(err)
	}
	// Vérification que la connection à la BDD marche
	pingErr := db.Ping()
	if pingErr != nil {
		println("2")
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Création du handler de la page d'accueil

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		// check if password is equal to passwordo
		if r.FormValue("password") != r.FormValue("passwordo") {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {

			albID, err := addUser(User{
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
	})
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}

func addUser(use User) (int64, error) {
	// Création de la requête SQL

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(use.Password), 8)
	result, err := db.Exec("INSERT INTO users (Username, Password, Passwordo, email) VALUES (?, ?, ?, ?)", use.Username, hashedPassword, use.Passwordo, use.Email)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	// Récupération de l'ID pour le return
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
