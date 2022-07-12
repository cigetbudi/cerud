package main

import (
	"cerud/controllers"
	"cerud/database"
	"fmt"
	"log"
	"net/http"

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
	log.Println(fmt.Println("Mulai jalankan server di port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), r))
}

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
