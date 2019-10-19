package main

var port string = ":5000"

func main() {
	// Create the mux router with the provided routes.
	router := CreateRouter()

	// Start the server specifying the port and the router.
	StartServer(port, *router)
}