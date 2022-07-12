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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var ps []entities.Product
	database.Debe.Find(&ps)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ps)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	pId := mux.Vars(r)["id"]
	if !checkIfProductExists(pId) {
		json.NewEncoder(w).Encode("Produk tidak ditemukan")
		return
	}
	var p entities.Product
	database.Debe.First(&p, pId)
	json.NewDecoder(r.Body).Decode(&p)
	database.Debe.Save(&p)
	w.Header().Set("COntent-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pId := mux.Vars(r)["id"]
	if !checkIfProductExists(pId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Produk tidak dtemukan")
		return
	}
	var p entities.Product
	database.Debe.Delete(&p, pId)
	json.NewEncoder(w).Encode("Produk berhasil dihapus")
}
