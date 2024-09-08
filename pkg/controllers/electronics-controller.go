package controllers

import (
	"alexandra1044/electronics-store-manager/pkg/models"
	"alexandra1044/electronics-store-manager/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewProduct models.Product

func GetProduct(w http.ResponseWriter, r *http.Request) {
	newProducts := models.GetAllProducts()
	res, _ := json.Marshal(newProducts)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	productDetails, _ := models.GetProductById(ID)
	res, _ := json.Marshal(productDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	CreateProduct := &models.Product{}
	utils.ParseBody(r, CreateProduct)
	p := CreateProduct.CreateProduct()
	res, _ := json.Marshal(p)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]
	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	product := models.DeleteProduct(ID)
	res, _ := json.Marshal(product)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	UpdateProduct := &models.Product{}
	utils.ParseBody(r, UpdateProduct)
	vars := mux.Vars(r)
	productId := vars["productId"]

	ID, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}

	productDetails, db := models.GetProductById(ID)
	if UpdateProduct.Name != "" {
		productDetails.Name = UpdateProduct.Name
	}
	if UpdateProduct.Price != 0 {
		productDetails.Price = UpdateProduct.Price
	}
	if UpdateProduct.Category != "" {
		productDetails.Category = UpdateProduct.Category
	}
	db.Save(&productDetails)
	res, _ := json.Marshal(productDetails)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
