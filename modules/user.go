package modules

import(
	"ttapi/models"
	"ttapi/tools"
//	"log"
)

var UserModule *User

type User struct {}  

func init(){
	UserModule = new(User)
}

func (u *User)CreateUser(user *models.User) bool {
	return tools.PH.CreateObject(user)
}
func (u *User)SelectUserByName(name string) models.User {
	return tools.PH.SelectObjectByFilter("name = "+ name, &models.User{}).(models.User)
}

