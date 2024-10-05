package routes

import (
	"net/http"
	"webapp/controllers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB, prefix string) {
	r.HandleFunc(prefix+"/user/create", func(w http.ResponseWriter, req *http.Request) {
		controllers.CreateUser(w, req, db)
	}).Methods("POST")
}
