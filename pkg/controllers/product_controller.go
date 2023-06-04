package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "project_demo/pkg/models/model_products"
	"project_demo/pkg/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var NewProduct models.Products

func GetProduct(w http.ResponseWriter, r *http.Request) {
	NewProduct := models.GetAllProducts()
	res, _ := json.Marshal(NewProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id_product"]
	ID, err := uuid.Parse(productID)
	if err != nil {
		fmt.Println("error while parsing", err)
	}

	productDetail, _ := models.GetProductById(ID.String())
	res, _ := json.Marshal(productDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	createProduct := &models.Products{}
	utils.ParseBody(r, createProduct)
	b := createProduct.Createproduct()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["id_product"]
	ID, err := uuid.Parse(productId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product := models.DeleteProduct(ID.String())
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var UpdateProduct = &models.Products{}
	utils.ParseBody(r, UpdateProduct)
	vars := mux.Vars(r)
	productId := vars["id_product"]
	ID, err := uuid.Parse(productId)

	if err != nil {
		fmt.Println("error while parsing")
	}
	productDetails, db := models.GetProductById(ID.String())

	if UpdateProduct.Name != "" {
		productDetails.Name = UpdateProduct.Name
	}
	if UpdateProduct.Price != 0 {
		productDetails.Price = UpdateProduct.Price
	}
	if UpdateProduct.Image != "" {
		productDetails.Image = UpdateProduct.Image
	}
	if UpdateProduct.Amount != 0 {
		productDetails.Amount = UpdateProduct.Amount
	}
	if UpdateProduct.Sale != 0 {
		productDetails.Sale = UpdateProduct.Sale
	}
	if UpdateProduct.Set != "" {
		productDetails.Set = UpdateProduct.Set
	}
	db.Save(&productDetails)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
