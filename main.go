package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	typedefs "./typedefs"
	utils "./utils"
	"io/ioutil"
	"net/http"
)

func ServePostList(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON.
	utils.SetContentType(&w)

	// Enable CORS on the server.
	utils.EnableCors(&w)

	// URL to fetch the content from.
	url := "https://www.hotelflow.com.br/wp-json/wp/v2/posts"

	// Makes a GET request to the provided URL.
	response, err := http.Get(url)
	utils.LogError(err)

	// Read the response.
	responseJSON, readErr := ioutil.ReadAll(response.Body)
	utils.LogError(readErr)

	// Initialize the struct variable to pass the JSON response as the content by reference.
	var postsList []typedefs.Posts

	// Parse the JSON response.
	jsonErr := json.Unmarshal(responseJSON, &postsList)
	utils.LogError(jsonErr)

	// Encode the JSON.
	json.NewEncoder(w).Encode(postsList)
}

func main() {
	// Create the main mux router.
	router := mux.NewRouter()

	// Create the routes.
	router.HandleFunc("/posts", ServePostList).Methods("GET")

	// Serve the routes.
	http.ListenAndServe(":5000", router)
}