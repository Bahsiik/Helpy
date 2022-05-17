package main

import (
	"fmt"
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

// It selects all the subjects from the database, and then for each subject it selects the topic name and the user name
// from the topics and users tables respectively
func selectAllSubjects() []Subject {
	stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date, Topic_id, User_id FROM subjects")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
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

func selectSubjectByTopic(topicID int) []Subject {
	stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date, Topic_id, User_id FROM subjects WHERE Topic_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicID)
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
	fmt.Println("SELECTED Subjects:")
	for _, subject := range subjects {
		fmt.Println("ID:", subject.ID, ", NAME:", subject.Name, ", DATE:", subject.Date, ", TOPIC ID:", subject.TopicID, ", USER ID:", subject.UserID)
	}
	return subjects
}

func selectSubjectTopicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		topicID := r.FormValue("topicID")
		if topicID == "" {
			fmt.Fprintf(w, "Please fill in all fields")
			return
		}
		stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date, Topic_id, User_id FROM subjects WHERE Topic_id = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		rows, err := stmt.Query(topicID)
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
		fmt.Println("SELECTED Subjects:")
		for _, subject := range subjects {
			fmt.Println("ID:", subject.ID, ", NAME:", subject.Name, ", DATE:", subject.Date, ", TOPIC ID:", subject.TopicID, ", USER ID:", subject.UserID)
		}
		// get cookie
		c, err := r.Cookie("session")
		if err != nil {
			fmt.Println("No c found")
		}
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
}
