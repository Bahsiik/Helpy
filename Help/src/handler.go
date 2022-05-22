package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var TMPL *template.Template

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterHandler ****")
	err := TMPL.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}

func RegisterAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** RegisterAuthHandler ****")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	username := r.FormValue("username")
	firstPassword := r.FormValue("firstPassword")
	secondPassword := r.FormValue("secondPassword")
	nameAlphaNumeric, nameLength := CheckUsername(username)
	if firstPassword != secondPassword {
		fmt.Println("**** Passwords don't match ****")
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Les mots de passe ne correspondent pas")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	passwordLowercase, passwordUppercase, passwordNumber, passwordSpecial, passwordLength, passwordNoSpaces := CheckPassword(firstPassword)
	if !passwordLowercase || !passwordUppercase || !passwordNumber || !passwordSpecial || !passwordLength || !passwordNoSpaces || !nameAlphaNumeric || !nameLength {
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Veuillez vérifier les critères de nom d'utilisateur ou de votre mot de passe")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	stmt := "SELECT Username FROM users WHERE Username = ?"
	row := DB.QueryRow(stmt, username)
	var uID string
	err = row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Ce nom d'utilisateur existe déjà")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(firstPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	var insertStmt *sql.Stmt
	insertStmt, err = DB.Prepare("INSERT INTO users (Username, Password) VALUES (?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	defer insertStmt.Close()
	_, err = insertStmt.Exec(username, hash)
	if err != nil {
		fmt.Println("error inserting new user:", err)
		err = TMPL.ExecuteTemplate(w, "register.html", "⚠ Il y a eu une erreur lors de la création du compte")
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		return
	}
	fmt.Fprint(w, "congrats, your account has been successfully created")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Handler ****")
	err := TMPL.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** Login Auth Handler ****")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	var hash string
	err = DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hash)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du hash du mot de passe de l'utilisateur: ", err)
		done := LoginError(w, err, "Veuillez vérifier vos identifiants")
		if done {
			return
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		done := LoginError(w, err, "Veuillez vérifier vos identifiants")
		if done {
			return
		}
		return
	} else {
		AddSessionCookie(w, r)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** indexHandler ****")
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**** homeHandler ****")
	d := GetUserInfoFromSession(w, r)
	S := SelectAllPost()
	d.Posts = S
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
		d.Posts[i].Hour = TranslateHour(d.Posts[i].RawDate)
	}
	err := TMPL.ExecuteTemplate(w, "home.html", d)
	if err != nil {
		return
	}
}

func SelectPostTopicHandler(w http.ResponseWriter, r *http.Request) {
	d := GetUsernameFromSession(w, r)
	r.ParseForm()
	topicID := r.FormValue("topicID")
	topicID = TranslateTopicID(topicID)
	postList := SelectPostByTopic(topicID)
	d.Posts = postList
	for i := 0; i < len(d.Posts); i++ {
		d.Posts[i].Date = TranslateDate(d.Posts[i].RawDate)
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
	topicID = TranslateTopicID(topicID)
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
	for i := 0; i < len(d.Replies); i++ {
		d.Replies[i].ReplyDate = TranslateDate(d.Replies[i].ReplyRawDate)
	}
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
	Content := SelectContentFromReplyID(FirstReplyID)
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
	ReplyContent := SelectContentFromReplyID(ReplyID)
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
	d.FirstPost.Content = SelectContentFromReplyID(d.ReplyID)
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

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** aboutHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "about.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** teamHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "team.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SettingProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingProfileHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingProfile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SettingNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingNotificationsHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingNotifications.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SettingAccountHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** settingAccountHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "settingAccount.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profile.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileComHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileComHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileCom.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileFavHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileFavHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileFav.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfileLikeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profileLikeHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profileLike.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func ProfilePubHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** profilePubHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "profilePub.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ThematiquesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** thematiquesHandler ***")
	d := GetUsernameFromSession(w, r)
	err := TMPL.ExecuteTemplate(w, "thematiques.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
