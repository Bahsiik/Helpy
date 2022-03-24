package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))

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
		log.Fatal(err)
	}
	// Vérification que la connection à la BDD marche
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//  Recherche des albums d'un artiste défini
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Recherche d'album via l'ID
	alb, err := albumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	//Suppression d'un album
	albTitle, err := delAlbum(Album{Title: "Deux Frères"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(albTitle)

	// Gestion de la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "index.gohtml", "")
		if err != nil {
			return
		}
		//name := r.FormValue("userinput")
		//print(name)
		prix, err := strconv.ParseFloat(r.FormValue("prix"), 32)
		// Ajout d'un album
		albID, err := addAlbum(Album{
			Title:  r.FormValue("titre"),
			Artist: r.FormValue("artiste"),
			Price:  float32(prix),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID of added album: %v\n", albID)
	})

	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		return
	}
}

// albumsByArtist recherche des albums d'un artiste défini
func albumsByArtist(name string) ([]Album, error) {
	// Variable pour contenir les données renvoyées par la fonction
	var albums []Album

	// Création de la requête SQL
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	// Permet de relâcher les informations contenues dans la variable rows lorsque la fonction s'arrête
	defer rows.Close()
	// Boucle pour remplir le tableau []Album des données de la recherche
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
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
func addAlbum(alb Album) (int64, error) {

	// Création de la requête SQL
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
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
