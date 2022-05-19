package main

import (
	"fmt"
	"net/http"
	"time"
)

type Reply struct {
	ID        int
	Message   string
	ReplyDate *time.Time
	ReplyNbr  int
	PostID    int
	ReplyToID int
	UserID    int
}

func postFeedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postFeedHandler ***")
	cookie := checkCookie(w, r)
	userID := getUserIdFromSession(cookie.Value)
	username := getUsernameFromID(userID)

	err := r.ParseForm()
	if err != nil {
		return
	}
	ID := r.FormValue("PostID")
	replies := selectReplyFromPostID(ID)
	fmt.Println("Reply list :")
	for _, reply := range replies {
		fmt.Println(reply)
	}
	post := selectPostByID(ID)
	userName := getUsernameFromID(post.UserID)
	post.UserName = userName
	d := data{
		Username:  username,
		FirstPost: post,
		Replies:   replies,
	}
	err = tmpl.ExecuteTemplate(w, "postFeed.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func selectReplyFromPostID(postID string) []Reply {
	fmt.Println("*** selectReplyFromPostID ***")
	var reply []Reply
	rows, err := db.Query("SELECT * FROM replies WHERE Post_id = ?", postID)
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

func selectPostByID(postID string) Post {
	fmt.Println("*** selectPostByID ***")
	var p Post
	rows, err := db.Query("SELECT * FROM post WHERE Post_id = ?", postID)
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
