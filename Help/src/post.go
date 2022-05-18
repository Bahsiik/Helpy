package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Post struct {
	ID        int
	Name      string
	Date      *time.Time
	TopicID   int
	TopicName string
	UserID    int
	UserName  string
}

type PostError struct {
	Title   string
	Content string
	Topic   string
}

func selectAllPost() []Post {
	stmt, err := db.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func selectPostByTopic(topicID string) []Post {
	stmt, err := db.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post WHERE Topic_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicID)
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func getPost(rows *sql.Rows, err error) []Post {
	var postList []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Name, &post.Date, &post.TopicID, &post.UserID)
		if err != nil {
			log.Fatal(err)
		}
		postList = append(postList, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return postList
}

func getPostAttributs(postList []Post) {
	for i := range postList {
		stmt, err := db.Prepare("SELECT Topic_name FROM topics WHERE Topic_id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(postList[i].TopicID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err := rows.Scan(&postList[i].TopicName)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
		// inner join with users to get the user name
		stmt, err = db.Prepare("SELECT Username FROM users WHERE User_id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err = stmt.Query(postList[i].UserID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err := rows.Scan(&postList[i].UserName)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func selectPostTopicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	topicID := r.FormValue("topicID")
	topicID = translateTopicID(topicID)
	postList := selectPostByTopic(topicID)
	cookie, err := r.Cookie("session")
	userID := getUserIdFromSession(cookie.Value)
	username := getUsernameFromID(userID)
	data := data{
		Name:  username,
		Posts: postList,
	}
	err = tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		return
	}
}

func translateTopicID(topicID string) string {
	switch topicID {
	case "Marketing & Communication":
		return "1"
	case "Audiovisuel":
		return "2"
	case "Création & Digital Design":
		return "3"
	case "Architecture d'intérieur":
		return "4"
	case "Informatique":
		return "5"
	case "Web Management":
		return "6"
	case "3D, Animation & Jeux-vidéo":
		return "7"
	case "2D & Illustration Digitale":
		return "8"
	case "Campus Life":
		return "9"
	case "Administration":
		return "10"
	}
	return ""
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** postHandler ***")
	// get the cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	// get the user id from the cookie
	userID := getUserIdFromSession(cookie.Value)
	// get the username from the user id
	username := getUsernameFromID(userID)
	d := data{
		Name: username,
	}
	err = tmpl.ExecuteTemplate(w, "create.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addPostHandler ***")
	cookie, _ := r.Cookie("session")
	userID := getUserIdFromSession(cookie.Value)
	var postError PostError
	var checkError bool
	r.ParseForm()
	title := r.FormValue("title")
	if title == "" {
		postError.Title = "Veuillez entrer un titre"
		checkError = true
	}
	topicID := r.FormValue("category")
	if topicID == "" {
		postError.Topic = "Veuillez choisir une catégorie"
		checkError = true
	}
	description := r.FormValue("description")
	if description == "" {
		postError.Content = "Veuillez entrer une description"
		checkError = true
	}
	if checkError == true {
		d := data{
			Name:         getUsernameFromID(userID),
			AddPostError: postError,
		}
		fmt.Println("data : ", d)
		err := tmpl.ExecuteTemplate(w, "create.html", d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		addPost(title, description, topicID, userID)
		http.Redirect(w, r, "/index", http.StatusFound)
	}
}

func addPost(title string, description string, topicID string, userID int) {
	stmt, err := db.Prepare("INSERT INTO post (Title, Content, Topic_id, User_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, description, topicID, userID)
	if err != nil {
		log.Fatal(err)
	}
}
