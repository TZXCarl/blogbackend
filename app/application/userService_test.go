package application

import (
	"blogbackend/app/db"
	"fmt"
	"testing"
)

func Test_AddUser(t *testing.T) {
	db.Init()
	AddUser()
}

func Test_FindUserByNameAndPwdv(t *testing.T) {
	db.Init()
	user := FindUserByNameAndPwd("test", "123456")
	fmt.Println(user)
}

func Test_UpdateUser(t *testing.T) {
	db.Init()
	UpdateUser("test", "111111", "123456")
}

func Test_DeleteUser(t *testing.T) {
	db.Init()
	DeleteUser("5a23bc7bdd3acefd8d9a8e56")
}
