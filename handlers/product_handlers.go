package handlers

import (
	"api-products/models"
	"api-products/storage"
	"encoding/json"
	"net/http"
	"strconv"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := storage.GetAll()
	if err != nil {
		http.Error(w, "Erro ao buscar os produtos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	product, err = storage.Create(product)
	if err != nil {
		http.Error(w, "Erro ao criar o produto", http.StatusInternalServerError)
		return
	}

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

	product, err := storage.GetById(id)

	if err != nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
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

	updateProduct, found, err := storage.Update(id, updateProduct)
	if err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProduct)
	return

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	found, err := storage.Delete(id)
	if err != nil {
		http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
