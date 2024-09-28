package routes

import (
	"net/http"
	"webapp/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterProductRoute(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/product/create", func(w http.ResponseWriter, rq *http.Request) {
		controllers.CreateProduct(w, rq, db)
	}).Methods("POST")

	r.HandleFunc("/products", func(w http.ResponseWriter, rq *http.Request) {
		controllers.List(w, rq, db)
	}).Methods("GET")
}
