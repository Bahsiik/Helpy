package main

import (
	"fmt"
	"time"
)

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
		rows.Scan(&r.ID, &r.Message, &r.ReplyDate, &r.PostID, &r.ReplyToID, &r.ReplyUserID)

		if err != nil {
			fmt.Println(err)
			return reply
		}
		reply = append(reply, r)
	}
	return reply
}

// SelectRepliesByPostIDString It takes a post id as a string, and returns a slice of replies
func SelectRepliesByPostIDString(postID string) []Reply {
	fmt.Println("*** SelectRepliesByPostIDString ***")
	var reply []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return reply
	}
	defer rows.Close()
	for rows.Next() {
		var r Reply
		err := rows.Scan(&r.ID, &r.Message, &r.ReplyRawDate, &r.PostID, &r.ReplyToID, &r.ReplyUserID, &r.Deleted)
		if err != nil {
			fmt.Println(err)
			return reply
		}
		r.UserName = SelectUsernameFromID(r.ReplyUserID)
		reply = append(reply, r)
	}
	if len(reply) == 0 {
		return reply
	}
	reply = reply[1:]
	return reply
}

func SelectRepliesByPostIDInt(postID int) []Reply {
	fmt.Println("*** SelectRepliesByPostIDString ***")
	var reply []Reply
	rows, err := DB.Query("SELECT * FROM replies WHERE Post_id = ?", postID)
	if err != nil {
		fmt.Println(err)
		return reply
	}
	defer rows.Close()
	for rows.Next() {
		var r Reply
		err := rows.Scan(&r.ID, &r.Message, &r.ReplyRawDate, &r.PostID, &r.ReplyToID, &r.ReplyUserID, &r.Deleted)
		if err != nil {
			fmt.Println(err)
			return reply
		}
		r.UserName = SelectUsernameFromID(r.ReplyUserID)
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

func SelectReplyContentFromReplyID(replyID int) string {
	fmt.Println("*** SelectReplyContentFromReplyID ***")
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

func SelectReplyDateFromReplyID(replyID int) *time.Time {
	fmt.Println("*** SelectReplyDateFromReplyID ***")
	var replyDate *time.Time
	rows, err := DB.Query("SELECT reply_date FROM replies WHERE Reply_id = ?", replyID)
	if err != nil {
		fmt.Println(err)
		return replyDate
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&replyDate)
		if err != nil {
			fmt.Println(err)
			return replyDate
		}
	}
	return replyDate
}
