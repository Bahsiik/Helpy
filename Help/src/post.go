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
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.RawDate, &post.ReplyNbr, &post.TopicID, &post.PostUserID)
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
		rows, err = stmt.Query(postList[i].PostUserID)
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

func TranslateTopicNameToTopicID(topicID string) string {
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

func TranslateTopicNameToTopicShortName(topicID string) string {
	switch topicID {
	case "Marketing & Communication":
		return "MarketCom"
	case "Audiovisuel":
		return "Audiovisuel"
	case "Création & Digital Design":
		return "CreaDesign"
	case "Architecture d'intérieur":
		return "Architecture"
	case "Informatique":
		return "Info"
	case "Web Management":
		return "WebMgmt"
	case "3D, Animation & Jeux-vidéo":
		return "Anim"
	case "2D & Illustration Digitale":
		return "Illu"
	case "Campus Life":
		return "CampusLife"
	case "Administration":
		return "Admin"
	}
	return ""
}

func DeletePost(postID string) {
	fmt.Println("DeletePost")
	stmt, err := DB.Prepare("DELETE FROM post WHERE Post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(postID)
	if err != nil {
		log.Fatal(err)
	}
}

func PostErrorRedirect(w http.ResponseWriter, r *http.Request, postError PostError) bool {
	d := GetUserInfoFromSession(w, r)
	d.AddPostError = postError
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

type Notification struct {
	UserID     int
	Notif_name string
	Notif_id   int
}

func notifPost(postID string) {
	fmt.Println("notifPost")
	stmt, err := DB.Prepare("SELECT User_id FROM post WHERE Post_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(postID)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			log.Fatal(err)
		}
		notif := Notification{
			UserID:     userID,
			Notif_name: "post",
			Notif_id:   1,
		}
		InsertNotification(notif)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func InsertNotification(notif Notification) {
	fmt.Println("InsertNotification")
	stmt, err := DB.Prepare("INSERT INTO notifications (User_id, Notif_type, Notif_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(notif.UserID, notif.Notif_name, notif.Notif_id)
	if err != nil {
		log.Fatal(err)
	}

}
