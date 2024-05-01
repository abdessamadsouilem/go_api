package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-playground/validator/v10"
)
// functions to validate the request 
var Validate = validator.New()
// function to parse request to json 
func ParseJson(r *http.Request,payload any) error {
	if r.Body == nil {
		fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

//function to write the json to return it 
func WriteJson(w http.ResponseWriter,status int,v any) error{
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// function to write error 
func WriteError(w http.ResponseWriter,status int,err error) {
	WriteJson(w,status,map[string]string{"error":err.Error()})
}