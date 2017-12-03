package application

import (
	"blogbackend/app/db"
	"blogbackend/app/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// func FindUserByNameAndPwd() models.User {
// 	user := models.User{}
// 	db.GetByQ(db.User, bson.M{"username": "cral", "password": "12121"}, &user)
// 	return user
// }

func InsertNote(note models.Note) bool {
	fmt.Println("insert....")
	return db.Insert(db.Note, note)
}

func findNote(userId, noteId string) models.Note {
	note := models.Note{}
	db.GetByQ(db.Note, bson.M{"userId": userId, "_id": bson.ObjectIdHex(noteId)}, &note)
	// db.Get2(db.Note, bson.ObjectId(noteId), &note)
	return note
}

func findNotes(userId string) []*models.Note {
	// fmt.Println(userId)
	list := []*models.Note{}
	db.ListByQ(db.Note, bson.M{"userId": userId}, &list)
	return list
}
