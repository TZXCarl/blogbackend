package application

import (
	"blogbackend/app/db"
	"blogbackend/app/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func AddUser() {
	user := models.NewUser("test", "zxtang", "123456", true)
	db.Insert(db.User, user)
}

func FindUserByNameAndPwd(username, password string) models.User {
	user := models.User{}
	db.GetByQ(db.User, bson.M{"username": username, "password": password}, &user)
	return user
}

func UpdateUser(username, password, newpassword string) bool {

	user := FindUserByNameAndPwd(username, password)
	oldPassword := user.Password
	user.Password = newpassword
	res := db.UpdateByQMap(db.User, bson.M{"username": user.Username, "password": oldPassword}, &user)
	fmt.Println(res)
	return res
}

func DeleteUser(id string) bool {
	user := models.User{}
	db.Get(db.User, id, &user)

	if user.Name == "" {
		panic("user is not exist!")
	}
	fmt.Println(id)
	res := db.Delete(db.User, bson.M{"_id": bson.ObjectIdHex(id)})
	if res == false {
		panic("failed")
	}
	return true
}
