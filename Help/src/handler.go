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
	d := GetUsernameFromSession(w, r)
	S := SelectAllPost()
	d.Posts = S
	fmt.Println("d: ", d)
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
	d = Data{Posts: postList}
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
	description := r.FormValue("description")
	postError, checkError := CheckPostError(title, description, topicID)
	if checkError == true {
		if PostErrorRedirect(w, userID, postError) {
			return
		}
	} else {
		AddPost(title, description, topicID, userID)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func PostFeedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postFeedHandler ***")
	d := GetUsernameFromSession(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}
	PostName := r.FormValue("PostName")
	post := SelectPostByName(PostName)
	userName := SelectUsernameFromID(post.UserID)
	post.UserName = userName
	d.FirstPost = post
	postID := SelectPostIDByName(PostName)
	replies := SelectRepliesByPostName(postID)
	d.Replies = replies
	fmt.Println("replies: ", replies)
	err = TMPL.ExecuteTemplate(w, "postFeed.html", d)
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
