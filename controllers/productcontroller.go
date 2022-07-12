package controllers

import (
	"cerud/database"
	"cerud/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p entities.Product
	json.NewDecoder(r.Body).Decode(&p)
	database.Debe.Create(&p)
	json.NewEncoder(w).Encode(p)

}

func checkIfProductExists(productId string) bool {
	var p entities.Product
	database.Debe.First(&p, productId)
	if p.ID == 0 {
		return false
	}
	return true
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	pId := mux.Vars(r)["id"]
	if !checkIfProductExists(pId) {
		json.NewEncoder(w).Encode("Product tidak ditemukan")
		return
	}
	var p entities.Product
	database.Debe.First(&p, pId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

}
