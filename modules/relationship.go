package modules

import(
	"ttapi/models"
	"ttapi/tools"
	"strconv"
	"log"
)

var RelationshipModule *Relationship

type Relationship struct {}  

func init(){
	RelationshipModule = new(Relationship)
}

func (r *Relationship)CreateRelationship(relationship *models.Relationship) bool {
	db := tools.PH.SelectObjectByFilter()
	defer db.Close()
	_, err := db.Exec("SELECT relationship_" + relationship.State + "(" + strconv.Itoa(relationship.User_id) +
		"," + strconv.Itoa(relationship.Other_user_id) + ");")
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
// func (r *Relationship)SelecRelationship(name string) []models.Relationship {
// 	var users []models.User
// 	db := tools.PH.SelectObjectByFilter()
// 	defer db.Close()
// 	err := db.Model(&models.User{}).Select(&users)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return users
// }

