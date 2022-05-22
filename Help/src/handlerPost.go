package main

import (
	"fmt"
	"net/http"
)

func SelectPostTopicHandler(w http.ResponseWriter, r *http.Request) {
	d := GetUsernameFromSession(w, r)
	r.ParseForm()
	topicName := r.FormValue("topicID")
	d.TopicShortName = TranslateTopicNameToTopicShortName(topicName)
	d.Topic = topicName
	topicID := TranslateTopicNameToTopicID(topicName)
	postList := SelectPostByTopic(topicID)
	d.Posts = postList
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
	}
	err := TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}

func SortPostHandler(w http.ResponseWriter, r *http.Request) {
	d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	sort := r.FormValue("sortType")
	if sort == "Date ⬆️" {
		d.Posts = SelectPostByDateUp()
	} else if sort == "Date ⬇️" {
		d.Posts = SelectPostByDateDown()
	} else if sort == "Popularité ⬆️" {
		d.Posts = SelectPostByRepliesUp()
	} else if sort == "Popularité ⬇️" {
		d.Posts = SelectPostByRepliesDown()
	}
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
	}
	err := TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "create.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addPostHandler ***")
	cookie := CheckCookie(w, r)
	userID := SelectUserIDFromSessionID(cookie.Value)
	r.ParseForm()
	title := r.FormValue("title")
	topicID := r.FormValue("category")
	topicID = TranslateTopicNameToTopicID(topicID)
	description := r.FormValue("description")
	postError, checkError := CheckPostError(title, description, topicID)
	if checkError == true {
		if PostErrorRedirect(w, userID, postError) {
			return
		}
	} else {
		AddPost(title, description, topicID, userID)
		postID := SelectPostIDByTitle(title)
		AddFirstReply(description, userID, postID)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** deletePostHandler ***")
	//d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	postID := r.FormValue("postID")
	DeleteRepliesFromPostID(postID)
	DeletePost(postID)
	http.Redirect(w, r, "/index", http.StatusFound)
}

func DeletePostAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** deletePostAdminHandler ***")
	//d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	postID := r.FormValue("postID")
	DeleteRepliesFromPostID(postID)
	DeletePost(postID)
	http.Redirect(w, r, "/admin", http.StatusFound)
}

func PostFeedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postFeedHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	PostName := r.FormValue("PostName")
	d.Replies = SelectRepliesByPostIDString(SelectPostIDByName(PostName))
	d.FirstPost = SelectPostByName(PostName)
	d.FirstPost.UserName = SelectUsernameFromID(d.FirstPost.PostUserID)
	d.FirstPost.Date = TranslateDate(d.FirstPost.RawDate)
	d.FirstPost.Hour = TranslateHour(d.FirstPost.RawDate)
	for i := 0; i < len(d.Replies); i++ {
		d.Replies[i].ReplyDate = TranslateDate(d.Replies[i].ReplyRawDate)
		d.Replies[i].ReplyHour = TranslateHour(d.Replies[i].ReplyRawDate)
	}
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SearchPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** searchPostHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("search")
	post := SelectPostBySearch(value)
	d.Posts = post
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
	}
	getPostAttributs(d.Posts)
	err = TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
