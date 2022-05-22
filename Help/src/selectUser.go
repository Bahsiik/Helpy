package main

import (
	"fmt"
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

func SelectAdminFromID(userID int) bool {
	var admin bool
	err := DB.QueryRow("SELECT Admin FROM users WHERE User_id = ?", userID).Scan(&admin)
	if err != nil {
		fmt.Println("err: ", err)
		return false
	}
	return admin
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