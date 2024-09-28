package main

import (
	"log"
	"net/http"
	"webapp/migrations"
	"webapp/routes"

	"github.com/gorilla/mux"
)

func main() {

	dsn := "root:PythonisSucks98@#@tcp(127.0.0.1:3306)/webapp_go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := migrations.Migrate(dsn)

	if err != nil {
		panic(err)
	}

	log.Println("Application started successfully")
	log.Println("Server running on localhost:8080")

	r := mux.NewRouter()

	routes.RegisterRoutes(r, db)

	log.Fatal(http.ListenAndServe(":8080", r))

}
