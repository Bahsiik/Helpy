package main

import "time"

type Data struct {
	UserID       int
	Username     string
	FirstPost    Post
	Posts        []Post
	PostID       int
	AddPostError PostError
	Replies      []Reply
	Reply        Reply
	ReplyID      int
	Search       string
	Type         ReplyType
}

type Session struct {
	Token  string
	UserID int
	Date   *time.Time
}

type Post struct {
	ID         int
	Title      string
	Content    string
	RawDate    *time.Time
	Date       string
	ReplyNbr   int
	TopicID    int
	TopicName  string
	PostUserID int
	UserName   string
}

type PostError struct {
	Title   string
	Content string
	Topic   string
}

type Reply struct {
	ID           int
	Message      string
	ReplyRawDate *time.Time
	ReplyDate    string
	ReplyNbr     int
	PostID       int
	ReplyToID    int
	ReplyUserID  int
	Deleted      bool
	UserName     string
}

type ReplyType struct {
	Post  bool
	Reply bool
}
