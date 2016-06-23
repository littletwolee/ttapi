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
func (r *Relationship)GetRelationshipById(user_id int) []models.Relationship {
	var relationships []models.Relationship
	db := tools.PH.SelectObjectByFilter()
	defer db.Close()
	err := db.Model(&models.Relationship{}).Where("user_id = ?", user_id).Select(&relationships)
	if err != nil {
		log.Println(err)
	}
	return relationships
}

