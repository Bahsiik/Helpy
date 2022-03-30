package main

import "fmt"

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