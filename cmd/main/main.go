package main

import (
	"log"
	"net/http"

	// "project_demo/pkg/controllers"
	"project_demo/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/products/product_women", controllers.CreateProduct).Methods("POST", "OPTIONS", "DELETE", "PUT")

	handler := cors.AllowAll().Handler(r)

	routes.RegisterProductsRoutes(r)
	routes.RegisterAccountRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
