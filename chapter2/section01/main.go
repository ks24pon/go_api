package main

import (
	"time"
	"fmt"
)

// Comment構造体の定義
type Comment struct {
	CommentID int
	ArticleID int
	Message string
	// 日時を表す
	CreateAt time.Time
}

// 記事構造体の定義
type Article struct {
	ID int
	Title string
	Contents string
	UserName string
	NiceNum int
	CommentList []Comment
	CreateAt time.Time
}

// 構造体を実際に確かめる

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "test comment1",
		CreateAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "second comment",
		CreateAt: time.Now(),
	}
	article := Article{
		ID: 1,
		Title: "first article",
		Contents: "This is the test article.",
		UserName: "saki",
		NiceNum: 1,
		CommentList: []Comment{comment1, comment2},
		CreateAt: time.Now(),
	}
	fmt.Printf("%+v\n", article)
}
