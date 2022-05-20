package main

import "time"

type Data struct {
	Username     string
	FirstPost    Post
	Posts        []Post
	AddPostError PostError
	Replies      []Reply
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
	Date      *time.Time
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
}
