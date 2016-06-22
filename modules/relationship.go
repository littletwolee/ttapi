package modules

import(
	"ttapi/models"
	"ttapi/tools"
//	"log"
)

var RelationshipModule *Relationship

type Relationship struct {}  

func init(){
	RelationshipModule = new(Relationship)
}

func (r *Relationship)CreateRelationship(relationship *models.Relationship) bool {
	relationship.Type = "relationship"
	return tools.PH.CreateObject(relationship)
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

