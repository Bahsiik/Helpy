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
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, reply_number, Topic_id, User_id FROM post")
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
	stmt, err := DB.Prepare("SELECT Post_id, Title, creation_date, reply_number, Topic_id, User_id FROM post WHERE Topic_id = ?")
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
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.PostID, &r.ReplyToID, &r.UserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		reply = append(reply, r)
	}
	return reply
}

// SelectPostIDByName It takes a post name as a string, and returns the post ID as a string
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

// SelectRepliesByPostName It takes a post name as a string, and returns a slice of replies
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
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.PostID, &r.ReplyToID, &r.UserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		r.UserName = SelectUsernameFromID(r.UserID)
		reply = append(reply, r)
	}
	if len(reply) == 0 {
		return reply
	}
	reply = reply[1:]
	return reply
}

func SelectFirstReplyIDByPostID(postID string) int {
	fmt.Println("*** SelectFirstReplyIDByPostID ***")
	var replyID int
	rows, err := DB.Query("SELECT Reply_id FROM replies WHERE Post_id = ? ORDER BY Reply_id ASC LIMIT 1", postID)
	if err != nil {
		fmt.Println(err)
		return replyID
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&replyID)
		if err != nil {
			fmt.Println(err)
			return replyID
		}
	}
	return replyID
}

func SelectPostIDFromStringID(postID string) int {
	fmt.Println("*** SelectPostIDFromStringID ***")
	var postID2 int
	rows, err := DB.Query("SELECT Post_id FROM replies WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return postID2
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&postID2)
		if err != nil {
			fmt.Println(err)
			return postID2
		}
	}
	return postID2
}

func SelectUsernameFromPostID(postID int) string {
	fmt.Println("*** SelectUsernameFromPostID ***")
	var username string
	// inner join to get username
	rows, err := DB.Query("SELECT username FROM users INNER JOIN post ON users.User_id = post.user_id WHERE post.Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return username
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
			return username
		}
	}
	return username
}

func SelectContentFromReplyID(replyID int) string {
	fmt.Println("*** SelectContentFromReplyID ***")
	var content string
	rows, err := DB.Query("SELECT Content FROM replies WHERE Reply_id = ?", replyID)
	if err != nil {
		fmt.Println(err)
		return content
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&content)
		if err != nil {
			fmt.Println(err)
			return content
		}
	}
	return content
}

func SelectPostIDByReplyID(replyID string) int {
	fmt.Println("*** SelectPostIDByReplyID ***")
	var postID int
	rows, err := DB.Query("SELECT Post_id FROM replies WHERE Reply_id = ?", replyID)
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

func SelectReplyIDFromStringID(replyID string) int {
	fmt.Println("*** SelectReplyIDFromStringID ***")
	var replyID2 int
	rows, err := DB.Query("SELECT Reply_id FROM replies WHERE Reply_id = ?", replyID)
	if err != nil {
		fmt.Println(err)
		return replyID2
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&replyID2)
		if err != nil {
			fmt.Println(err)
			return replyID2
		}
	}
	return replyID2
}

func SelectUsernameFromReplyID(replyID int) string {
	fmt.Println("*** SelectUsernameFromReplyID ***")
	var username string
	rows, err := DB.Query("SELECT username FROM users INNER JOIN replies ON users.User_id = replies.User_id WHERE Reply_id = ?", replyID)
	if err != nil {
		fmt.Println(err)
		return username
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
			return username
		}
	}
	return username
}
