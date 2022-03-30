package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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

func newAlbum(r *http.Request) {
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
}
