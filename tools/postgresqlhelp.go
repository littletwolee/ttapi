package tools

import (
	"gopkg.in/pg.v4"
	"log"
	"reflect"
//	"ttapi/models"
)

var (
	PH *PostgreSQLHelp
	postgresqlhost string
	postgresqlport string
	postgresqluser string
	postgresqlpwd string
	postgresqldbname string
	options *pg.Options
)

type PostgreSQLHelp struct{}

func init() {
	PH = new(PostgreSQLHelp)
	options = new(pg.Options)
	options.Addr = Conf.GetValue("postgresql", "postgresqlhost")
	options.Database = Conf.GetValue("postgresql", "postgresqldbname")
	options.User = Conf.GetValue("postgresql", "postgresqluser")
	options.Password = Conf.GetValue("postgresql", "postgresqlpwd")
}

func connect() *pg.DB {
	return pg.Connect(options)
}


func (p *PostgreSQLHelp) CreateObject(object interface{}) bool {
	db := connect()
	defer db.Close()
	err := db.Create(object)
	if err != nil {
		log.Println(err)
		return false 
	}
	return true
}

// func (p *PostgreSQLHelp) DeleteObjectByName() bool {
// 	db := connect()
// 	defer db.Close()
// 	err := db.Create(object)
// 	if err != nil {
// 		log.Println(err)
// 		return false 
// 	}
// 	return true
// }

func (p *PostgreSQLHelp) SelectObjectByFilter(filter string, object interface{}) interface{} {
	db := connect()
	defer db.Close()
	v := reflect.ValueOf(object)
	err := db.Model(&v).Where("name = ?", "IT").Select()
	if err != nil {
		log.Println(err)
		return nil
	}
	return object
}
