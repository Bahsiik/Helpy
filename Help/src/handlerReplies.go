package main

import (
	"fmt"
	"net/http"
)

func ReplyToPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** replyToPostHandler ***")
	d := GetUserInfoFromSession(w, r)
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
	d := GetUserInfoFromSession(w, r)
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
	d := GetUserInfoFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	value := r.FormValue("ReplyID")
	d.ReplyID = SelectReplyIDFromStringID(value)
	d.FirstPost.Content = SelectReplyContentFromReplyID(d.ReplyID)
	postID := SelectPostIDByReplyID(value)
	d.FirstPost.UserName = SelectUsernameFromPostID(postID)
	replyContent := r.FormValue("content")
	postError, checkError := CheckReplyError(replyContent)
	if postError == true {
		ReplyErrorRedirect(w, d, checkError)
	} else {
		AddReply(replyContent, d.UserID, postID, d.ReplyID)
		AddReplyNumberToPost(postID)
		d = GetPostFeedFromInt(d, postID)
		err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
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
	d.ReplyID = SelectReplyIDFromStringID(value)
	DeleteReplyFromReplyID(d.ReplyID)
	UpdateReplyStatus(d.ReplyID)
	postID := SelectPostIDByReplyID(value)
	RemoveReplyNumberFromPost(postID)

	d = GetPostFeedFromInt(d, postID)
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

	d = GetPostFeedFromInt(d, postID)
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
}

func GetPostFeedFromInt(d Data, postID int) Data {
	d.Replies = SelectRepliesByPostIDInt(postID)
	d.FirstPost = SelectPostByPostIDInt(postID)
	d.FirstPost.UserName = SelectUsernameFromID(d.FirstPost.PostUserID)
	d.FirstPost.Date = TranslateDate(d.FirstPost.RawDate)
	d.FirstPost.Hour = TranslateHour(d.FirstPost.RawDate)
	d.FirstPost.UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.FirstPost.UserName))
	for i := 0; i < len(d.Replies); i++ {
		d.Replies[i].ReplyDate = TranslateDate(d.Replies[i].ReplyRawDate)
		d.Replies[i].ReplyHour = TranslateHour(d.Replies[i].ReplyRawDate)
		d.Replies[i].RepliedMsgUserName = SelectUsernameFromReplyID(d.Replies[i].ReplyToID)
		d.Replies[i].RepliedMsgContent = SelectReplyContentFromReplyID(d.Replies[i].ReplyToID)
		d.Replies[i].RepliedMsgRawDate = SelectReplyDateFromReplyID(d.Replies[i].ID)
		d.Replies[i].RepliedMsgDate = TranslateDate(d.Replies[i].RepliedMsgRawDate)
		d.Replies[i].RepliedMsgHour = TranslateHour(d.Replies[i].RepliedMsgRawDate)
		d.Replies[i].UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.Replies[i].UserName))
	}
	return d
}
