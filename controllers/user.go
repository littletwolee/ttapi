package controllers

import (
	"ttapi/models"
	"ttapi/modules"
	"ttapi/tools"
	"encoding/json"
//	"fmt"
	"net/http"
//	"log"
	"io/ioutil"
	"strconv"
//	"github.com/gorilla/mux"
)

// Operations about object
type UserController struct {}

// @Title CreateUser
// @Description create user
// @Param	name
// @Success 200 {result} models.Result
// @Failure 500 
// @router /users[POST]
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	result := &models.Result{}
	r.ParseForm()
	body, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var user *models.User
	json.Unmarshal(body, &user)
	flag := modules.UserModule.CreateUser(user)
	result.StatusCode = http.StatusOK
	result.Data = map[string]string{"state":strconv.FormatBool(flag)}
	tools.RH.GetResult(w, result)
}


// @Title GetAllUsers
// @Description get all users
// @Success 200 {result} models.Result
// @Failure 500 
// @router /users[GET]
func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result := &models.Result{}
	list := modules.UserModule.GetAllUsers()
	result.StatusCode = http.StatusOK
	result.Data = list
	tools.RH.GetResult(w, result)
}
