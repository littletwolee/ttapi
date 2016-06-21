package main

//简单的JSON Restful API演示(服务端)
//author: Xiong Chuan Liang
//date: 2015-2-28

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"ttapi/tools"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	m := Message{"IT"}  
	resp, err := json.MarshalIndent(m, "", "")  
	if err != nil {  
		panic(err)  
	}  
	fmt.Fprintf(w, string(resp))  
}

func main() {
	conf := ConfigHelper.SetConfig("./conf/conf.ini")
	conf.SetConfig("ttapi/conf/app.conf")
	post := conf.Read("default", "httpport")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe("localhost:"+post, nil)
}

type Message struct {  
    Dept    string  
}  

