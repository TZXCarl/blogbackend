package models

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
	Age      string `json:"age"`
	Name     string `json:"name"`
}
