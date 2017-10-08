package controllers

import (
	"blogbackend/app/db"
	"github.com/revel/revel"
	
	"blogbackend/app/models"	
)

type App struct {
	*revel.Controller
}

// type User struct {
// 	name string
// 	Username string
// 	Password string
// 	IsAdmin bool
// }


func (c App) Index() revel.Result {

	var user models.User
	user.Name = "hello"
	user.Password = "12121"
	user.Username = "cral"
	user.IsAdmin = true

	// host, _ := revel.Config.String("db.host")
	db.Insert(db.User, user)

	// fmt.Println(host)
	test := "hello"
	return c.Render(test)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}