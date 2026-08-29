package main

import (
	"api-products/handlers"
	"api-products/storage"
	"fmt"
	"net/http"
)

func main() {
	storage.InitDB()

	http.HandleFunc("GET /products", handlers.ListProducts)
	http.HandleFunc("POST /products", handlers.CreateProduct)
	http.HandleFunc("GET /products/{id}", handlers.GetProduct)
	http.HandleFunc("PUT /products/{id}", handlers.UpdateProduct)
	http.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)

	fmt.Println("Server is running in port 8000")
	http.ListenAndServe(":8000", nil)
}
