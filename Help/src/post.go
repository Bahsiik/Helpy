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

//func addSubjectHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		r.ParseForm()
//		subjectName := r.FormValue("subjectName")
//		topicID := r.FormValue("topicID")
//		userID := r.FormValue("userID")
//		if subjectName == "" || topicID == "" || userID == "" {
//			fmt.Fprintf(w, "Please fill in all fields")
//			return
//		}
//		stmt, err := db.Prepare("INSERT INTO subjects (Subject_name, Topic_id, User_id) VALUES (?, ?, ?)")
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer stmt.Close()
//		_, err = stmt.Exec(subjectName, topicID, userID)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Fprintf(w, "Post added")
//	} else {
//		fmt.Fprintf(w, "Please use POST")
//	}
//}

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
		Name:     username,
		Subjects: postList,
	}
	err = tmpl.ExecuteTemplate(w, "post.html", data)
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
	case "Développement Web":
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
	err := tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** addPostHandler ***")
	r.ParseForm()
	title := r.FormValue("title")
	topicID := r.FormValue("category")
	description := r.FormValue("description")
	// get the user id from the session
	cookie, _ := r.Cookie("session")
	userID := getUserIdFromSession(cookie.Value)
	fmt.Println("*** addPostHandler ***")
	fmt.Println("Titre du nouveau : ", title)
	fmt.Println("ID du topic : ", topicID)
	fmt.Println("ID du user :", userID)
	addPost(title, description, topicID, userID)
	http.Redirect(w, r, "/index", http.StatusFound)
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
