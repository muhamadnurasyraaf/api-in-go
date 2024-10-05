package main

import (
	"flag"
	"log"
	"net/http"
	"webapp/database"
	"webapp/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:PythonisSucks98@#@tcp(127.0.0.1:3306)/webapp_go?charset=utf8mb4&parseTime=True&loc=Local"

	dropFlag := flag.Bool("drop", false, "Drop and recreate all tables")
	flag.Parse()

	var db *gorm.DB
	var err error

	if *dropFlag {
		db, err = database.MigrateFresh(dsn)
	} else {
		db, err = database.Migrate(dsn)
	}

	if err != nil {
		panic(err)
	}

	log.Println("Application started successfully")
	log.Println("Server running on localhost:8080")

	r := mux.NewRouter()
	routes.RegisterRoutes(r, db)

	routes.PrintRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))

}
