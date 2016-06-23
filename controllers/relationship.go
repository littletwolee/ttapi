package controllers

import (
	"ttapi/models"
	"ttapi/modules"
	"encoding/json"
	"fmt"
	"net/http"
//	"log"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)

// Operations about object
type RelationshipController struct {}

// @Title CreateUser
// @Description find user by objectid
// @Param	objectId	"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (rs *RelationshipController) CreateRelationship(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var relationship *models.Relationship 
	json.Unmarshal(body, &relationship)
	query := mux.Vars(r)
	user_id, err := strconv.Atoi(query["user_id"])
	if err != nil {
	}
	relationship.User_id = user_id
	other_user_id, err := strconv.Atoi(query["other_user_id"])
	if err != nil {
	}
	relationship.Other_user_id = other_user_id
	if relationship.State == "liked" || relationship.State == "unliked" {
		result := modules.RelationshipModule.CreateRelationship(relationship)
		resp, err := json.MarshalIndent(map[string]string{"msg":strconv.FormatBool(result)}, "" , "")
		if err != nil {  
			panic(err)  
		}
		fmt.Fprintf(w, string(resp))
	} else {

	}
	fmt.Fprintf(w, string(""))  
}

// func (r *RelationshipController) SelectUserByName(w http.ResponseWriter, r *http.Request) {
// 	query := mux.Vars(r)
// 	result := modules.UserModule.SelectUserByName(query["name"])
// 	resp, err := json.MarshalIndent(result, "" , "")
// 	if err != nil {  
// 		panic(err)  
// 	}  
// 	fmt.Fprintf(w, string(resp))  
// }
