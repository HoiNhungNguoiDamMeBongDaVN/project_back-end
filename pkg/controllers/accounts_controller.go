package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	models "project_demo/pkg/models/model_accounts"
	"project_demo/pkg/utils"

	// "strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var NewAccount models.Accounts

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	NewAccount := models.GetAllAccounts()
	res, _ := json.Marshal(NewAccount)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetProductWomenById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	productID := vars["productId"]
// 	ID, err := strconv.ParseComplex(productID, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing", err)
// 	}

// 	productDetail, _ := models.GetProductWomenById(ID)
// 	res, _ := json.Marshal(productDetail)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

func GetAccountById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["Id_account"]
	ID, err := uuid.Parse(accountId)
	if err != nil {
		fmt.Println("error while parsing", err)
	}

	productDetail, _ := models.GetAccountById(ID.String())
	res, _ := json.Marshal(productDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	createProduct := &models.Accounts{}
	utils.ParseBody(r, createProduct)
	b := createProduct.CreateAccount()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["Id_account"]
	ID, err := uuid.Parse(accountId)
	if err != nil {
		fmt.Println("error while parsing")
	}
	product := models.DeleteAccount(ID.String())
	res, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var UpdateAccont = &models.Accounts{}
	utils.ParseBody(r, UpdateAccont)
	vars := mux.Vars(r)
	accountId := vars["Id_account"]
	ID, err := uuid.Parse(accountId)

	if err != nil {
		fmt.Println("error while parsing")
	}
	accountDetails, db := models.GetAccountById(ID.String())

	if UpdateAccont.AdminName != "" {
		accountDetails.AdminName = UpdateAccont.AdminName
	}
	if UpdateAccont.PasswordAdmin != "" {
		accountDetails.PasswordAdmin = UpdateAccont.PasswordAdmin
	}
	db.Save(&accountDetails)
	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
