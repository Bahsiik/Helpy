package main

import (
	"fmt"
	"log"
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

func selectSubjects() []Subject {
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
	// inner join with topics to get the topic name
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

	fmt.Println("Subjects:")
	for _, subject := range subjects {
		fmt.Println("ID:", subject.ID, ", NAME:", subject.Name, ", DATE:", subject.Date, ", TOPIC ID:", subject.TopicID, ", USER ID:", subject.UserID, ", TOPIC NAME:", subject.TopicName, ", USER NAME:", subject.UserName)
	}
	return subjects
}
