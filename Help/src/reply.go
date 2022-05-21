package main

import (
	"fmt"
	"log"
	"net/http"
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

func AddReplyNumberToPost(postID int) {
	stmt, err := DB.Prepare("UPDATE post SET reply_number = reply_number + 1 WHERE Post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(postID)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckReplyError(content string) (bool, PostError) {
	var ReplyError PostError
	if content == "" {
		ReplyError.Content = "Vous n'avez rien écrit"
		return true, ReplyError
	}
	if len(content) > 1000 {
		ReplyError.Content = "Votre réponse est trop longue"
		return true, ReplyError
	}
	return false, ReplyError
}

func ReplyErrorRedirect(w http.ResponseWriter, d Data, ReplyError PostError) {
	d.AddPostError = ReplyError
	err := TMPL.ExecuteTemplate(w, "reply.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteRepliesFromPostID(postID string) {
	fmt.Println("Delete replies from postID: ", postID)
	stmt, err := DB.Prepare("DELETE FROM replies WHERE Post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(postID)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteReplyFromReplyID(replyID string) {
	fmt.Println("Delete reply from replyID: ", replyID)
	stmt, err := DB.Prepare("DELETE FROM replies WHERE Reply_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(replyID)
	if err != nil {
		log.Fatal(err)
	}
}
