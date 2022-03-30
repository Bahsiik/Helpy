package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

// Ça je ne sais pas trop, mais c'est ce qui nous permet de manip la bdd
var db *sql.DB

// Album Structure qui permettra d'interagir avec la bdd, reprend les memes attributs
type User struct {
	USERNAME string
	PASSWORD string
}

func main() {

	// Gestion de tous les fichiers gohtml
	tmpl := template.Must(template.ParseGlob("templates/login.gohtml"))
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

	////  Recherche des albums d'un artiste défini
	//albums, err := albumsByArtist("John Coltrane")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Albums found: %v\n", albums)
	//
	//// Recherche d'album via l'ID
	//alb, err := albumByID(1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Album found: %v\n", alb)
	//
	////Suppression d'un album
	//albTitle, err := delAlbum(Album{Title: "Deux Frères"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Name of the deleted album: %v\n", albTitle)
	//Gestion de la page d'accueil

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "login.gohtml", "")
		if err != nil {
			return
		}
		switch r.Method {
		case "POST":
			// Ajout d'un User
			albID, err := addUser(User{
				USERNAME: r.FormValue("Username"),
				PASSWORD: r.FormValue("password"),
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID of added album: %v\n", albID)
		}
	})
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}

func addUser(use User) (int64, error) {
	// Création de la requête SQL
	result, err := db.Exec("INSERT INTO users (Username, Password) VALUES (?, ?)", use.USERNAME, use.PASSWORD)
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
