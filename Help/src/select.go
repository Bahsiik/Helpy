package main

import (
	"fmt"
	"log"
)

// SelectUserIDFromSessionID It takes a sessionID as a string, and returns the userID as an int
func SelectUserIDFromSessionID(sessionID string) int {
	stmt := "SELECT User_id FROM session WHERE Session_id = ?"
	rows, err := DB.Query(stmt, sessionID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

// SelectUserIDFromUsername It takes a username as a string, and returns the userId as an int
func SelectUserIDFromUsername(username string) int {
	stmt := "SELECT User_id FROM users WHERE Username = ?"
	rows, err := DB.Query(stmt, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var userId int
	for rows.Next() {
		err := rows.Scan(&userId)
		if err != nil {
			panic(err)
		}
	}
	return userId
}

// SelectUsernameFromID This function takes in a userID and returns the username associated with that userID
func SelectUsernameFromID(userID int) string {
	stmt := "SELECT Username FROM users WHERE User_id = ?"
	rows, err := DB.Query(stmt, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
	}
	return username
}

// SelectAllPost We prepare a query to select all the posts from the database, then we execute the query and get the rows. We then call
// the getPost function to get the postList and then we call the getPostAttributs function to get the post attributs
func SelectAllPost() []Post {
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

// SelectPostByTopic It selects all the posts from the database that have the same topic ID as the one passed in the function
func SelectPostByTopic(topicID string) []Post {
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, Topic_id, User_id FROM post WHERE Topic_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(topicID)
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

// SelectPostByID It selects a post from the database by its ID
func SelectPostByID(postID string) Post {
	fmt.Println("*** SelectPostByID ***")
	var p Post
	rows, err := DB.Query("SELECT * FROM post WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Date, &p.ReplyNbr, &p.TopicID, &p.UserID)
		if err != nil {
			return Post{}
		}
		if err != nil {
			fmt.Println(err)
			return p
		}
	}
	return p
}

func SelectPostByName(postName string) Post {
	fmt.Println("*** SelectPostByName ***")
	var p Post
	rows, err := DB.Query("SELECT * FROM post WHERE Title = ?", postName)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Date, &p.ReplyNbr, &p.TopicID, &p.UserID)
		if err != nil {
			return Post{}
		}
		if err != nil {
			fmt.Println(err)
			return p
		}
	}
	return p
}

// SelectReplyFromPostID It takes a postID as a string, and returns a slice of Reply structs
func SelectReplyFromPostID(postID string) []Reply {
	fmt.Println("*** SelectReplyFromPostID ***")
	var reply []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return reply
	}
	defer rows.Close()
	for rows.Next() {
		var r Reply
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.ReplyNbr, &r.PostID, &r.ReplyToID, &r.UserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		reply = append(reply, r)
	}
	return reply
}

func SelectPostIDByName(postName string) string {
	fmt.Println("*** SelectPostIDByName ***")
	var postID string
	rows, err := DB.Query("SELECT Post_id FROM post WHERE Title = ?", postName)
	if err != nil {
		fmt.Println(err)
		return postID
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&postID)
		if err != nil {
			fmt.Println(err)
			return postID
		}
	}
	return postID
}

func SelectRepliesByPostName(postName string) []Reply {
	fmt.Println("*** SelectRepliesByPostName ***")
	var reply []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE Post_id = ?", postName)
	if err != nil {
		fmt.Println(err)
		return reply
	}
	defer rows.Close()
	for rows.Next() {
		var r Reply
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.ReplyNbr, &r.PostID, &r.ReplyToID, &r.UserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		r.UserName = SelectUsernameFromID(r.UserID)
		reply = append(reply, r)
	}
	return reply
}
