package main

import (
	"flag"
	"log"
	"net/http"
	"webapp/database"
	"webapp/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},         // Allow requests from port 3000
		AllowCredentials: true,                                      // Allow credentials (e.g., cookies, authorization headers)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Allowed HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed headers
	})

	handler := corsHandler.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))

}
