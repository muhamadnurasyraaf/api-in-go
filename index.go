package main

import (
	"log"
	"net/http"
	"webapp/migrations"

	"github.com/gorilla/mux"
)

func createProduct() {

}

func main() {

	dsn := "root:PythonisSucks98@#@tcp(127.0.0.1:3306)/webapp_go?charset=utf8mb4&parseTime=True&loc=Local"

	if err := migrations.Migrate(dsn); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Application started successfully")
	log.Println("Server running on localhost:8080")

	r := mux.NewRouter()

	r.HandleFunc("/product/create", createProduct()).Methods("POST")
	http.ListenAndServe(":8080", nil)

}
