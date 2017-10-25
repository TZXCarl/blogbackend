package application

import (
	"blogbackend/app/db"
	"blogbackend/app/models"

	"gopkg.in/mgo.v2/bson"
)

func FindUserByNameAndPwd(username, password string) models.User {
	user := models.User{}
	db.GetByQ(db.User, bson.M{"username": username, "password": password}, &user)
	return user
}
