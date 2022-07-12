package main

import (
	"cerud/controllers"
	"cerud/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	//load config.json pakai viper
	LoadConfig()

	//migrasi data
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	//init routes
	r := mux.NewRouter().StrictSlash(true)

	//regiter routes
	RegisterProductRoutes(r)

	//jalankan server

	port := os.Getenv("PORT")
	log.Println(fmt.Sprintf("Starting Server on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/api/product", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
