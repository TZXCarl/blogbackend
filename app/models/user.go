package models

import (
	"github.com/revel/revel"
	"fmt"
)

type User struct {
	Name string
	Username string
	Password string
	IsAdmin bool
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}


func (user *User) Validate(v *revel.Validation) {

}

// func (user *User) setName(name string){
// 	user.name = name
// }

// func (user *User) setUsername(username string){
// 	user.Username = username
// }

// func (user *User) setPassword(password string){
// 	user.Password = password
// }

// func (user *User) setIsAdmin(IsAdmin bool){
// 	user.IsAdmin = IsAdmin
// }