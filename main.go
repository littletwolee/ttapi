package main

import (
	"net/http"
	_ "ttapi/tools"
	_ "ttapi/routers"
	"ttapi/tools"
//	"log"
)

func main() {
	post := tools.Conf.GetValue("default", "httpport")
	http.ListenAndServe("localhost:" + post, nil)
}


