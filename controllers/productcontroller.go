package controllers

import (
	"cerud/database"
	"cerud/entities"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p entities.Product
	json.NewDecoder(r.Body).Decode(&p)
	database.Debe.Create(&p)
	json.NewEncoder(w).Encode(p)

}
