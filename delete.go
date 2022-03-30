package main

import (
	"fmt"
	"log"
)

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

func deleteAlbum(err error) {
	//Suppression d'un album
	albTitle, err := delAlbum(Album{Title: "Deux Frères"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name of the deleted album: %v\n", albTitle)
}
