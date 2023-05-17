package routes

import (
	"project_demo/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterAccountRoutes = func(router *mux.Router) {
	router.HandleFunc("/account", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account", controllers.GetAccounts).Methods("GET")
	router.HandleFunc("/account/{Id_account}", controllers.GetAccountById).Methods("GET")
	router.HandleFunc("/account/{Id_account}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{Id_account}", controllers.DeleteAccount).Methods("DELETE")
}
