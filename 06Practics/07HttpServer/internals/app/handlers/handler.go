package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func WrapError(w http.ResponseWriter, err error){
	WrapErrorWithStatus(w, err, http.StatusInternalServerError)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int){
	var m = map[string]string{
		"result": "error",
		"data": err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Xontext-Type-Options","nosniff")
	w.WriteHeader(httpStatus)
	fmt.Println(w, string(res))
}
