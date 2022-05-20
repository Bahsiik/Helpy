package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var DB *sql.DB

func AddPost(title string, description string, topicID string, userID int) {
	stmt, err := DB.Prepare("INSERT INTO post (Title, Content, Topic_id, User_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, description, topicID, userID)
	if err != nil {
		log.Fatal(err)
	}
}

// It takes a pointer to a sql.Rows object and an error, and returns a slice of Post objects
func getPost(rows *sql.Rows, err error) []Post {
	var postList []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.RawDate, &post.ReplyNbr, &post.TopicID, &post.UserID)
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

func SelectPostIDByTitle(title string) int {
	var postID int
	err := DB.QueryRow("SELECT Post_id FROM post WHERE Title = ?", title).Scan(&postID)
	if err != nil {
		log.Fatal(err)
	}
	return postID
}

// It takes a list of posts and for each post it gets the topic name and the user name and adds them to the post
func getPostAttributs(postList []Post) {
	for i := range postList {
		stmt, err := DB.Prepare("SELECT Topic_name FROM topics WHERE Topic_id = ?")
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
		stmt, err = DB.Prepare("SELECT Username FROM users WHERE User_id = ?")
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

func TranslateTopicID(topicID string) string {
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

func PostErrorRedirect(w http.ResponseWriter, userID int, postError PostError) bool {
	d := Data{
		Username:     SelectUsernameFromID(userID),
		AddPostError: postError,
	}
	fmt.Println("Data : ", d)
	err := TMPL.ExecuteTemplate(w, "create.html", d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func CheckPostError(title string, description string, topicID string) (PostError, bool) {
	var postError PostError
	var checkError bool
	if title == "" {
		postError.Title = "Veuillez entrer un titre"
		checkError = true
	}
	if description == "" {
		postError.Content = "Veuillez entrer une description"
		checkError = true
	}
	if topicID == "" {
		postError.Topic = "Veuillez choisir une catégorie"
		checkError = true
	}
	return postError, checkError
}
