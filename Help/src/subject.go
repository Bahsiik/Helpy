package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Subject struct {
	ID        int
	Name      string
	Date      *time.Time
	TopicID   int
	TopicName string
	UserID    int
	UserName  string
}

func selectAllSubjects() []Subject {
	stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date, Topic_id, User_id FROM subjects")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	subjects := getSubjects(rows, err)
	getSubjectAttributs(subjects)
	return subjects
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
//		fmt.Fprintf(w, "Subject added")
//	} else {
//		fmt.Fprintf(w, "Please use POST")
//	}
//}

func selectSubjectByTopic(topicID string) []Subject {
	stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date, Topic_id, User_id FROM subjects WHERE Topic_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicID)
	subjects := getSubjects(rows, err)
	getSubjectAttributs(subjects)
	return subjects
}

func getSubjects(rows *sql.Rows, err error) []Subject {
	var subjects []Subject
	for rows.Next() {
		var subject Subject
		err := rows.Scan(&subject.ID, &subject.Name, &subject.Date, &subject.TopicID, &subject.UserID)
		if err != nil {
			log.Fatal(err)
		}
		subjects = append(subjects, subject)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return subjects
}

func getSubjectAttributs(subjects []Subject) {
	for i := range subjects {
		stmt, err := db.Prepare("SELECT Topic_name FROM topics WHERE Topic_id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(subjects[i].TopicID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err := rows.Scan(&subjects[i].TopicName)
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
		rows, err = stmt.Query(subjects[i].UserID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err := rows.Scan(&subjects[i].UserName)
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

func selectSubjectTopicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	topicID := r.FormValue("topicID")
	topicID = translateTopicID(topicID)
	subjects := selectSubjectByTopic(topicID)
	c, err := r.Cookie("session")
	userID := getUserIdFromSession(c.Value)
	username := getUsernameFromID(userID)
	d := data{
		Name:     username,
		Subjects: subjects,
	}
	err = tmpl.ExecuteTemplate(w, "post.html", d)
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
