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
	d.Posts = SelectPostByUsername(d.Username)
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
		d.Posts[i].UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.Posts[i].UserName))
		fmt.Println(d.Posts[i].UserAvatar)
	}
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

func SelectPostByUsername(username string) []Post {
	fmt.Println("*** SelectPostByUsername ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post WHERE User_id = (SELECT User_id FROM users WHERE Username = ?)", username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var post Post
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.RawDate, &post.ReplyNbr, &post.TopicID, &post.PostUserID)
		posts = append(posts, post)
	}
	return posts
}

func SelectRepliesByUsername(username string) []Reply {
	fmt.Println("*** SelectRepliesByUsername ***")
	var replies []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE User_id = (SELECT User_id FROM users WHERE Username = ?)", username)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var reply Reply
		rows.Scan(&reply.ID, &reply.Message, &reply.ReplyRawDate, &reply.PostID, &reply.ReplyToID, &reply.ReplyUserID, &reply.Deleted)
		replies = append(replies, reply)
	}
	return replies
}
