package main

import (
	controllers "./controllers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router{
	// Create the main mux router.
	router := mux.NewRouter()

	// Create the routes.
	router.HandleFunc("/posts", controllers.ReturnParsedPostList).Methods("GET")

	return router
}