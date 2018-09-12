package controllers

import (
	"rest-api/services"

	"github.com/gorilla/mux"
)

type ControllerConfig struct {
	authService     services.AuthService
	contactsService services.ContactsService
}

func NewcontrollerConfig(as services.AuthService, cs services.ContactsService) *ControllerConfig {
	return &ControllerConfig{
		authService:     as,
		contactsService: cs,
	}
}

func InitializeRouter(router *mux.Router, cc *ControllerConfig) {

	router.Use(cc.authService.JwtAuthentication)
	router.HandleFunc("/contacts", cc.GetAllContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", cc.GetContact).Methods("GET")
	router.HandleFunc("/create-contact", cc.CreateContact).Methods("POST")


	router.HandleFunc("/create-user", cc.CreateUser).Methods("POST")
	router.HandleFunc("/login", cc.Login).Methods("POST")

}
