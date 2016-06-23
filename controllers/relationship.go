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

// @Title CreateRelationship
// @Description create relationship from user_id to other_user_id
// @Param	user_id	other_user_id state
// @Success 200 {result} models.Result
// @Failure 500 
// @router /users/{user_id}/relationships/{other_user_id}[PUT]
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
		tools.CHBool <- modules.RelationshipModule.CreateRelationship(relationship)
		flag := <- tools.CHBool
		result.StatusCode = http.StatusOK
		result.Data = map[string]string{"state":strconv.FormatBool(flag)}
		tools.RH.GetResult(w, result)
	} else {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter state error"
		tools.RH.GetResult(w, result)
	}
}


// @Title GetRelationshipById
// @Description get relationships by user_id
// @Param	user_id
// @Success 200 {result} models.Result
// @Failure 500 
// @router /users/{user_id}/relationships[GET]
func (u *RelationshipController) GetRelationshipById(w http.ResponseWriter, r *http.Request) {
	result := &models.Result{}
	query := mux.Vars(r)
	user_id, err := strconv.Atoi(query["user_id"])
	if err != nil {
		result.StatusCode = http.StatusBadRequest
		result.Errmsg = "parameter user_id error"
		tools.RH.GetResult(w, result)
	}
	tools.CHRelationship <- modules.RelationshipModule.GetRelationshipById(user_id)
	list := <- tools.CHRelationship
	result.StatusCode = http.StatusOK
	result.Data = list
	tools.RH.GetResult(w, result)
}

