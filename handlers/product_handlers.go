package handlers

import (
	"api-products/models"
	"api-products/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.Products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	product.ID = storage.NextID
	storage.NextID++

	storage.Products = append(storage.Products, product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, product := range storage.Products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var updateProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	for i, product := range storage.Products {
		if product.ID == id {
			updateProduct.ID = id

			storage.Products[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, product := range storage.Products {
		if product.ID == id {
			storage.Products = append(storage.Products[:i], storage.Products[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Produto não encontrado", http.StatusNotFound)
}
