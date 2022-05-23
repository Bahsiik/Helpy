package main

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileHandler ***")
	d := GetUserInfoFromSession(w, r)
	d.LastPost = SelectLastPost(d.UserID)
	d.LastPost.Date = TranslateDate(d.LastPost.RawDate)
	d.LastPost.Hour = TranslateHour(d.LastPost.RawDate)
	err := TMPL.ExecuteTemplate(w, "profile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SelectLastPost(userId int) Post {
	fmt.Println("*** SelectLastPost ***")
	// select the last post of the user with the given userId and return it
	var post Post
	DB.QueryRow("SELECT * FROM post WHERE User_id = ? ORDER BY creation_date DESC LIMIT 1", userId).Scan(&post.ID, &post.Title, &post.Content, &post.RawDate, &post.ReplyNbr, &post.TopicID, &post.PostUserID)
	return post
}
