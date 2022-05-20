package main

import (
	"fmt"
	"log"
)

func SelectAllPost() []Post {
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func SelectPostByTopic(topicID string) []Post {
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post WHERE Topic_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicID)
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func SelectReplyFromPostID(postID string) []Reply {
	fmt.Println("*** SelectReplyFromPostID ***")
	var reply []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return reply
	}
	defer rows.Close()
	for rows.Next() {
		var r Reply
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.ReplyNbr, &r.PostID, &r.ReplyToID, &r.UserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		reply = append(reply, r)
	}
	return reply
}

func SelectPostByID(postID string) Post {
	fmt.Println("*** SelectPostByID ***")
	var p Post
	rows, err := DB.Query("SELECT * FROM post WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Date, &p.ReplyNbr, &p.TopicID, &p.UserID)
		if err != nil {
			return Post{}
		}
		if err != nil {
			fmt.Println(err)
			return p
		}
	}
	return p
}

// SelectUserIdFromUsername Get the user id from the database from the username  (if the user exists)
func SelectUserIdFromUsername(username string) int {
	stmt := "SELECT User_id FROM users WHERE Username = ?"
	rows, err := DB.Query(stmt, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

func SelectUserIdFromSessionID(sessionID string) int {
	stmt := "SELECT User_id FROM session WHERE Session_id = ?"
	rows, err := DB.Query(stmt, sessionID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

func SelectUsernameFromID(userID int) string {
	stmt := "SELECT Username FROM users WHERE User_id = ?"
	rows, err := DB.Query(stmt, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
	}
	return username
}
