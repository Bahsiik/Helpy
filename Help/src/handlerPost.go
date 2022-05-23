package main

import (
	"fmt"
	"net/http"
)

func SelectPostTopicHandler(w http.ResponseWriter, r *http.Request) {
	d := GetUserInfoFromSession(w, r)
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
		d.Posts[i].UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.Posts[i].UserName))
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
		d.Topic = "Date ⬆️"
		d.TopicShortName = "all"
	} else if sort == "Date ⬇️" {
		d.Posts = SelectPostByDateDown()
		d.Topic = "Date ⬇️"
		d.TopicShortName = "all"
	} else if sort == "Popularité ⬆️" {
		d.Posts = SelectPostByRepliesUp()
		d.Topic = "Popularité ⬆️"
		d.TopicShortName = "all"
	} else if sort == "Popularité ⬇️" {
		d.Posts = SelectPostByRepliesDown()
		d.Topic = "Popularité ⬇️"
		d.TopicShortName = "all"
	}
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
		d.Posts[i].UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.Posts[i].UserName))
	}
	err := TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postHandler ***")
	d := GetUserInfoFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "create.html", d)
	fmt.Println(d.IsAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addPostHandler ***")
	d := GetUserInfoFromSession(w, r)
	r.ParseForm()
	title := r.FormValue("title")
	topicID := r.FormValue("category")
	topicID = TranslateTopicNameToTopicID(topicID)
	description := r.FormValue("description")
	postError, checkError := CheckPostError(title, description, topicID)
	if checkError == true {
		if PostErrorRedirect(w, r, postError) {
			return
		}
	} else {
		AddPost(title, description, topicID, d.UserID)
		postID := SelectPostIDByTitle(title)
		AddFirstReply(description, d.UserID, postID)
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
	d = GetPostFeedFromString(d, PostName)
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPostFeedFromString(d Data, PostName string) Data {
	d.Replies = SelectRepliesByPostIDString(SelectPostIDByName(PostName))
	d.FirstPost = SelectPostByName(PostName)
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
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
		d.Posts[i].UserAvatar = TranslateAvatarIdToString(SelectAvatarIdFromUsername(d.Posts[i].UserName))
	}
	getPostAttributs(d.Posts)
	err = TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
