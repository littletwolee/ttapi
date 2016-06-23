package controllers

import (
	"ttapi/models"
	"ttapi/modules"
	"encoding/json"
	"ttapi/tools"
//	"fmt"
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
	result := &models.Result{}
	r.ParseForm()
	body, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var relationship *models.Relationship 
	json.Unmarshal(body, &relationship)
	query := mux.Vars(r)
	user_id, err := strconv.Atoi(query["user_id"])
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter user_id error"
		tools.RH.GetResult(w, result)
	}
	relationship.User_id = user_id
	other_user_id, err := strconv.Atoi(query["other_user_id"])
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter other_user_id error"
		tools.RH.GetResult(w, result)
	}
	relationship.Other_user_id = other_user_id
	if relationship.State == "liked" || relationship.State == "unliked" {
		flag := modules.RelationshipModule.CreateRelationship(relationship)
		result.StatusCode = http.StatusOK
		result.Data = map[string]string{"state":strconv.FormatBool(flag)}
		tools.RH.GetResult(w, result)
	} else {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter state error"
		tools.RH.GetResult(w, result)
	}
}

func (u *RelationshipController) GetRelationshipById(w http.ResponseWriter, r *http.Request) {
	result := &models.Result{}
	query := mux.Vars(r)
	user_id, err := strconv.Atoi(query["user_id"])
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter user_id error"
		tools.RH.GetResult(w, result)
	}
	list := modules.RelationshipModule.GetRelationshipById(user_id)
	result.StatusCode = http.StatusOK
	result.Data = list
	tools.RH.GetResult(w, result)
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
