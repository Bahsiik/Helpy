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
	Type         ReplyType
}

type Session struct {
	Token  string
	UserID int
	Date   *time.Time
}

type Post struct {
	ID        int
	Title     string
	Content   string
	RawDate   *time.Time
	Date      string
	ReplyNbr  int
	TopicID   int
	TopicName string
	UserID    int
	UserName  string
}

type PostError struct {
	Title   string
	Content string
	Topic   string
}

type Reply struct {
	ID        int
	Message   string
	ReplyDate *time.Time
	ReplyNbr  int
	PostID    int
	ReplyToID int
	UserID    int
	UserName  string
}

type ReplyType struct {
	Post  bool
	Reply bool
}
