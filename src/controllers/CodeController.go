package controllers

import (
	"encoding/json"
	"models"
	"net/http"
	u "utils"
)

var CreateCode = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	code := &models.Code{}

	err := json.NewDecoder(r.Body).Decode(code)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	code.UserId = user
	resp := code.Create()
	u.Respond(w, resp)
}

var GetCodeFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetCodes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var RunCode = func(w http.ResponseWriter,r *http.Request){
	id := r.Context().Value("user").(uint)
	code := models.Code{}
	json.NewDecoder(r.Body).Decode(&code)
	code.UserId=id
	retVal,retCode:=code.Run()
	var msg string
	if retCode{
		msg ="Success"
	} else {
		msg = "Failed"
	}
	resp := u.Message(retCode,msg)
	resp["Output"] = retVal
	u.Respond(w, resp)
}