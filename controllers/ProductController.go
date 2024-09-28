package controllers

import (
	"encoding/json"
	"net/http"
	"webapp/models"
	"webapp/repositories"

	"gorm.io/gorm"
)

func CreateProduct(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	productRepo := repositories.NewProductRepository(db)

	if err := productRepo.Create(&product); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func List(w http.ResponseWriter, r *http.Request, db *gorm.DB) {

	productRepo := repositories.NewProductRepository(db)

	products, err := productRepo.GetAll()

	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
