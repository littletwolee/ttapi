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
type UserController struct {}

// @Title CreateUser
// @Description find user by objectid
// @Param	objectId	"the objectid you want to get"
// @Success 200 {user} models.User
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var user *models.User
	json.Unmarshal(body, &user)
	result := modules.UserModule.CreateUser(user)
	resp, err := json.MarshalIndent(map[string]string{"msg":strconv.FormatBool(result)}, "" , "")
	if err != nil {  
		panic(err)  
	}  
	fmt.Fprintf(w, string(resp))  
}

func (u *UserController) SelectUserByName(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	result := modules.UserModule.SelectUserByName(query["name"])
	resp, err := json.MarshalIndent(result, "" , "")
	if err != nil {  
		panic(err)  
	}  
	fmt.Fprintf(w, string(resp))  
}
