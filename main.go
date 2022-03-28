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
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))
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
			albID, err := addAlbum(User{
				USERNAME: r.FormValue("username"),
				PASSWORD: r.FormValue("password"),
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID of added mod: %v\n", albID)

		}
	})
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}

// albumsByArtist recherche des albums d'un artiste défini
func albumsByArtist(name string) ([]User, error) {
	// Variable pour contenir les données renvoyées par la fonction
	var User []User

	// Création de la requête SQL
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	// Permet de relâcher les informations contenues dans la variable rows lorsque la fonction s'arrête
	defer rows.Close()
	// Boucle pour remplir le tableau []Album des données de la recherche
	for rows.Next() {
		var Use User
		if err := rows.Scan(&Use.USERNAME, &Use.PASSWORD); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(User, Use)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return User, nil
}

// albumByID recherche d'album via l'ID.
func albumByID(id int64) (Album, error) {
	// Variable de type album pour contenir les données du retour de la recherche
	var alb Album

	// Création de la requête SQL
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil

}

// addAlbum Ajout d'un album,
// Renvoi également l'ID du nouvel album
func addAlbum(use User) (int64, error) {

	// Création de la requête SQL
	result, err := db.Exec("INSERT INTO album (Username, Password) VALUES (?, ?)", use.USERNAME, use.PASSWORD)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	// Récupération de l'ID pour le return
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// delAlbum suppression d'un album voulu,
// Renvoi également le nom de l'album supprimé
func delAlbum(alb Album) (string, error) {
	// Récupération du nom pour le return
	albTitle := alb.Title
	// Création de la requête SQL
	_, err := db.Exec("DELETE FROM album WHERE title=(?)", alb.Title)
	if err != nil {
		return "", fmt.Errorf("delAlbum: %v", err)
	}
	return albTitle, nil
}
func ModAlbum(alb Album) (int64, error) {

	result, err := db.Exec("UPDATE album SET id= ?, title= ?, artist= ?, price =? WHERE id = ?", alb.ID, alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return 0, fmt.Errorf("ModAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ModAlbum: %v", err)
	}
	return id, nil
}
