package main

import "github.com/gorilla/mux"

func main() {

}

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
