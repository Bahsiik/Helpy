package main

import (
	"fmt"
	"log"
)

// SelectAllPost We prepare a query to select all the posts from the database, then we execute the query and get the rows. We then call
// the getPost function to get the postList and then we call the getPostAttributs function to get the post attributs
func SelectAllPost() []Post {
	stmt, err := DB.Prepare("SELECT * FROM post")
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
	stmt, err := DB.Prepare("SELECT * FROM post WHERE Topic_id = ?")
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
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Date, &p.ReplyNbr, &p.TopicID, &p.PostUserID)
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
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.RawDate, &p.ReplyNbr, &p.TopicID, &p.PostUserID)
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

func SelectPostByPostIDString(postID string) Post {
	fmt.Println("*** SelectPostByPostIDString ***")
	var p Post
	rows, err := DB.Query("SELECT * FROM post WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.RawDate, &p.ReplyNbr, &p.TopicID, &p.PostUserID)
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

func SelectPostByPostIDInt(postID int) Post {
	fmt.Println("*** SelectPostByPostIDString ***")
	var p Post
	rows, err := DB.Query("SELECT * FROM post WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return p
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.RawDate, &p.ReplyNbr, &p.TopicID, &p.PostUserID)
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

func SelectPostIDByTitle(title string) int {
	var postID int
	err := DB.QueryRow("SELECT Post_id FROM post WHERE Title = ?", title).Scan(&postID)
	if err != nil {
		log.Fatal(err)
	}
	return postID
}

func SelectPostBySearch(search string) []Post {
	fmt.Println("*** SelectPostBySearch ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post WHERE Title LIKE ? OR Content LIKE ?", "%"+search+"%", "%"+search+"%")
	if err != nil {
		fmt.Println(err)
		return posts
	}
	defer rows.Close()
	postList := getPost(rows, err)
	return postList
}

func SelectPostByDateUp() []Post {
	fmt.Println("*** SelectPostByDateUp ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post ORDER BY creation_date DESC")
	if err != nil {
		fmt.Println(err)
		return posts
	}
	defer rows.Close()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func SelectPostByDateDown() []Post {
	fmt.Println("*** SelectPostByDateDown ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post ORDER BY creation_date ASC")
	if err != nil {
		fmt.Println(err)
		return posts
	}
	defer rows.Close()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func SelectPostByRepliesUp() []Post {
	fmt.Println("*** SelectPostByRepliesUp ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post ORDER BY reply_number DESC")
	if err != nil {
		fmt.Println(err)
		return posts
	}
	defer rows.Close()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}

func SelectPostByRepliesDown() []Post {
	fmt.Println("*** SelectPostByRepliesDown ***")
	var posts []Post
	rows, err := DB.Query("SELECT * FROM post ORDER BY reply_number ASC")
	if err != nil {
		fmt.Println(err)
		return posts
	}
	defer rows.Close()
	postList := getPost(rows, err)
	getPostAttributs(postList)
	return postList
}
