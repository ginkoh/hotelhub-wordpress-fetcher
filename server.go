package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer(port string, router mux.Router) {
	// Serve the routes.
	http.ListenAndServe(port, &router)
}