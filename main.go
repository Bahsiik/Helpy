package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

// Ça je ne sais pas trop, mais c'est ce qui nous permet de manip la bdd
var db *sql.DB

// Album Structure qui permettra d'interagir avec la bdd, reprend les memes attributs
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	// Gestion de tous les fichiers gohtml
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))

	// Paramètres de connexion à la BDD
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
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

	//deleteAlbum(err)
	//Gestion de la page d'accueil

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.gohtml", "")
		if err != nil {
			return
		}
		switch r.Method {
		case "GET":
			//////  Recherche des albums d'un artiste défini
			//albums, err := albumsByArtist("John Coltrane")
			//if err != nil {
			//	log.Fatal(err)
			//}
			//fmt.Printf("Albums found: %v\n", albums)
			tmpl.ExecuteTemplate(w, "index.gohtml", "")
		case "POST":
			searchByArtist(err, r)
			newAlbum(r)
		}
	})
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}
