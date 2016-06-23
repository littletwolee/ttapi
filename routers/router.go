// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ttapi/controllers"
	"github.com/gorilla/mux"
	"net/http"
//	"log"
)

func init() {
	r := mux.NewRouter()
	userController := new(controllers.UserController)
	user := "/users"
	r.HandleFunc(user, userController.CreateUser).
		Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc(user, userController.GetAllUsers).
		Methods("GET")
	relationshipController := new(controllers.RelationshipController)
	relationship := "/relationships"
	r.HandleFunc(user + "/{user_id}" + relationship + "/{other_user_id}", relationshipController.CreateRelationship).
		Methods("PUT").Headers("Content-Type", "application/json")
	r.HandleFunc(user + "/{user_id}" + relationship, relationshipController.GetRelationshipById).
		Methods("GET")
	http.Handle("/", r)
}
