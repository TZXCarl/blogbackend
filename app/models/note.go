package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Note struct {
	Id          bson.ObjectId `bson:"_id"`
	Title       string        `bson:"title"`
	Content     string        `bson:"content"`
	UserId      string        `bson:"userId"`
	CreatedTime int64         `bson:"createdTime"`
	UpdateTime  int64         `bson:"updateTime"`
}

func NewNote(title, content string) *Note {
	note := new(Note)
	// note.id = bson.NewObjectId()
	note.Title = title
	note.Content = content
	if note.Title == "" {
		fmt.Println("title id empty")
		return nil
	}
	if note.Content == "" {
		fmt.Println("content is empty")
		return nil
	}
	return note
}
