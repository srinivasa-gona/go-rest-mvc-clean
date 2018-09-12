package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (cc ControllerConfig) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	var people []models.Contact

	people, err := cc.contactsService.GetAllContacts()
	if err != nil {
		log.Fatalln("Error in get all response  ", err)
	}
	json.NewEncoder(w).Encode(people)
}

func (cc ControllerConfig) CreateContact(w http.ResponseWriter, r *http.Request) {
	var p models.Contact
	_ = json.NewDecoder(r.Body).Decode(&p)
	cc.contactsService.CreateContact(&p)
	json.NewEncoder(w).Encode(&p)

}

func (cc ControllerConfig) GetContact(w http.ResponseWriter, r *http.Request) {
	var people []models.Contact

	params := mux.Vars(r)
	for _, item := range people {
		itemId, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			return
		}
		if item.ID == itemId {
			json.NewEncoder(w).Encode(item)
		}
	}
	return

}
