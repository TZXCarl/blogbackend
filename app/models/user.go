package models

import (
	"fmt"

	"github.com/revel/revel"
)

type User struct {
	// id       bson.ObjectId `bson:"_id"`  //bson.ObjectId: 主键,   bson.NewObjectId():生成新的主键
	Name     string `bson:"name"`
	Username string `bson:"username"`
	Password string
	IsAdmin  bool
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

func NewUser(name string, username string, password string, isadmin bool) *User {
	ret := new(User)
	ret.Name = name
	ret.Username = username
	ret.Password = password
	ret.IsAdmin = isadmin
	if ret.Name == "" {
		fmt.Println("name is empty")
		return nil
	}

	if ret.Username == "" {
		fmt.Println("username is empty")
		return nil
	}
	if ret.Password == "" {
		fmt.Println("password is empty")
		return nil
	}
	return ret
}

func (user *User) Validate(v *revel.Validation) {

}
