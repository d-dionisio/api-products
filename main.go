package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Monitor", Price: 700},
	{ID: 2, Name: "Teclado", Price: 150},
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("GET /products", listProducts)

	fmt.Println("Server is running in port 8000")
	http.ListenAndServe(":8000", nil)
}
