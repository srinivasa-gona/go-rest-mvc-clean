package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
	"rest-api/utils"
)

func (cc ControllerConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	response, err := cc.authService.CreateUser(&u)
	utils.RestRespond(w, response, err)
}

func (cc ControllerConfig) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	response, err := cc.authService.Login(u.Username, u.Password)
	utils.RestRespond(w, response, err)
}
