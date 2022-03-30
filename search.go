package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

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

func searchByArtist(err error, r *http.Request) {
	artistName := r.FormValue("artistSearch")
	//  Recherche des albums d'un artiste défini
	albums, err := albumsByArtist(artistName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
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

func searchByID(err error) {
	// Recherche d'album via l'ID
	alb, err := albumByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
}
