package modules

import(
	"ttapi/models"
	"ttapi/tools"
	"log"
)

var UserModule *User

type User struct {}  

func init(){
	UserModule = new(User)
}

func (u *User)CreateUser(user *models.User) bool {
	user.Type = "user"
	return tools.PH.CreateObject(user)
}
func (u *User)GetAllUsers() []models.User {
	var users []models.User
	db := tools.PH.SelectObjectByFilter()
	defer db.Close()
	err := db.Model(&models.User{}).Select(&users)
	if err != nil {
		log.Println(err)
	}
	return users
}

