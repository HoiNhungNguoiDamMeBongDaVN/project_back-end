package routes

import (
	"project_demo/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterProductsRoutes = func(router *mux.Router) {
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id_product}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/products/{id_product}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id_product}", controllers.DeleteProduct).Methods("DELETE")
}
