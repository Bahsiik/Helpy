package main

import (
	"log"
)

func AddFirstReply(content string, userID int, postID int) {
	stmt, err := DB.Prepare("INSERT INTO replies (Content, User_id, Post_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, userID, postID)
	if err != nil {
		log.Fatal(err)
	}
}

func AddReply(content string, userID int, postID int, replyID int) {
	stmt, err := DB.Prepare("INSERT INTO replies (Content, User_id, Post_id, ReplyTo_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, userID, postID, replyID)
	if err != nil {
		log.Fatal(err)
	}
}
