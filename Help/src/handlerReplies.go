package main

import (
	"fmt"
	"net/http"
)

func ReplyToPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** replyToPostHandler ***")
	d := GetUsernameFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("Topic")
	PostID := SelectPostIDFromStringID(value)
	Username := SelectUsernameFromPostID(PostID)
	FirstReplyID := SelectFirstReplyIDByPostID(value)
	Content := SelectReplyContentFromReplyID(FirstReplyID)
	d.FirstPost.UserName = Username
	d.FirstPost.Content = Content
	d.ReplyID = FirstReplyID
	err = TMPL.ExecuteTemplate(w, "reply.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReplyToReplyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** replyToReplyHandler ***")
	d := GetUsernameFromSession(w, r)
	d = GetUserIDFromSession(w, r)
	d.Username = SelectUsernameFromID(d.UserID)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("Topic")
	ReplyID := SelectReplyIDFromStringID(value)
	ReplyContent := SelectReplyContentFromReplyID(ReplyID)
	Username := SelectUsernameFromReplyID(ReplyID)
	d.FirstPost.UserName = Username
	d.FirstPost.Content = ReplyContent
	d.ReplyID = ReplyID
	err = TMPL.ExecuteTemplate(w, "reply.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddReplyToPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addReplyToPostHandler ***")
	d := GetUsernameFromSession(w, r)
	d = GetUserIDFromSession(w, r)
	d.Username = SelectUsernameFromID(d.UserID)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("ReplyID")
	d.ReplyID = SelectReplyIDFromStringID(value)
	postID := SelectPostIDByReplyID(value)
	d.FirstPost.UserName = SelectUsernameFromPostID(postID)
	d.FirstPost.Content = SelectReplyContentFromReplyID(d.ReplyID)
	replyContent := r.FormValue("content")
	postError, checkError := CheckReplyError(replyContent)
	if postError == true {
		ReplyErrorRedirect(w, d, checkError)
	} else {
		AddReply(replyContent, d.UserID, postID, d.ReplyID)
		AddReplyNumberToPost(postID)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func DeleteReplyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** deleteReplyHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("ReplyID")
	fmt.Println("ReplyID: " + value)
	d.ReplyID = SelectReplyIDFromStringID(value)
	DeleteReplyFromReplyID(d.ReplyID)
	UpdateReplyStatus(d.ReplyID)
	postID := SelectPostIDByReplyID(value)
	RemoveReplyNumberFromPost(postID)

	d.Replies = SelectRepliesByPostIDInt(postID)
	d.FirstPost = SelectPostByPostIDInt(postID)
	d.FirstPost.UserName = SelectUsernameFromID(d.FirstPost.PostUserID)
	d.FirstPost.Date = TranslateDate(d.FirstPost.RawDate)
	for i := 0; i < len(d.Replies); i++ {
		d.Replies[i].ReplyDate = TranslateDate(d.Replies[i].ReplyRawDate)
	}
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
}

func DeleteReplyAdminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** deleteReplyAdminHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("ReplyID")
	fmt.Println("ReplyID: " + value)
	d.ReplyID = SelectReplyIDFromStringID(value)
	DeleteReplyFromReplyIDAdmin(d.ReplyID)
	UpdateReplyStatus(d.ReplyID)
	postID := SelectPostIDByReplyID(value)
	RemoveReplyNumberFromPost(postID)

	d.Replies = SelectRepliesByPostIDInt(postID)
	d.FirstPost = SelectPostByPostIDInt(postID)
	d.FirstPost.UserName = SelectUsernameFromID(d.FirstPost.PostUserID)
	d.FirstPost.Date = TranslateDate(d.FirstPost.RawDate)
	for i := 0; i < len(d.Replies); i++ {
		d.Replies[i].ReplyDate = TranslateDate(d.Replies[i].ReplyRawDate)
	}
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
}
