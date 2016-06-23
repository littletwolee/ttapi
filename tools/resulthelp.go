package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ttapi/models"
	"strconv"
)
var RH *ResultHelper

type ResultHelper struct{}

func init() {
	RH = new(ResultHelper)
}

func (rh *ResultHelper)GetResult(w http.ResponseWriter, r *models.Result){
	resp, err := json.MarshalIndent(r, "" , "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, string(
			"{\"StatusCode\":\"" + strconv.Itoa(http.StatusInternalServerError) + "\"," +
				"\"Errmsg\":\"Server Error\"," +
				"\"Data\":\"\"}"))
	}
	w.WriteHeader(r.StatusCode)
	fmt.Fprintf(w, string(resp))
}
