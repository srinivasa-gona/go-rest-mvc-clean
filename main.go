package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/controllers"
	"rest-api/repository"
	"rest-api/services"
	"rest-api/utils"

	//	"rest-api/utils"

	"github.com/gorilla/mux"
)

func main() {

	log.SetOutput(os.Stderr)
	router := mux.NewRouter()
	initializeApplication(router)
	log.Fatal(http.ListenAndServe(":8000", router))

}

func initializeApplication(router *mux.Router) {
	config := utils.GetConfiguration()
	log.Print(config.DB_USERNAME)
	db, _ := repository.GetCconnection(&config)
	authRepo := repository.NewAuthRepositoryImpl(db)
	contactsRepo := repository.NewContactsRepositoryImpl(db)
	authServ := services.NewAuthServiceImpl(authRepo, &config)
	contactServ := services.NewContactsServiceImpl(contactsRepo)
	controllerConfig := controllers.NewcontrollerConfig(authServ, contactServ)
	controllers.InitializeRouter(router, controllerConfig)
}
