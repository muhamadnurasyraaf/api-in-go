package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var prefix string = "/api"

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	RegisterProductRoute(r, db, prefix)
	RegisterUserRoutes(r, db, prefix)
}

func PrintRoutes(r *mux.Router) {
	fmt.Println("Available routes:")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate() // Get the path of the route
		if err != nil {
			return err
		}
		methods, _ := route.GetMethods() // Get the methods allowed for the route
		fmt.Printf("%s\t%s\n", methods, path)
		return nil
	})
}
