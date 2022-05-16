package main

import (
	"fmt"
	"log"
	"time"
)

type Subject struct {
	ID   int
	Name string
	Date *time.Time
}

func selectSubjects() []Subject {
	stmt, err := db.Prepare("SELECT Subject_id, Subject_name, creation_date FROM subjects")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	var subjects []Subject
	for rows.Next() {
		var subject Subject
		err := rows.Scan(&subject.ID, &subject.Name, &subject.Date)
		if err != nil {
			log.Fatal(err)
		}
		subjects = append(subjects, subject)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// print the subjects
	fmt.Println("Subjects:")
	for _, subject := range subjects {
		fmt.Println("ID:", subject.ID, ", NAME:", subject.Name, ", DATE:", subject.Date)
	}
	return subjects
}
