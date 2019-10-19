package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"fmt"
	typedefs "./typedefs"
	"io/ioutil"
	"log"
	"net/http"
)


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func ServePostList(w http.ResponseWriter, r *http.Request) {
	setContentType(&w)
	enableCors(&w)

	url := "https://www.hotelflow.com.br/wp-json/wp/v2/posts"

	// Creates a client connection.
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Read the response.
	responseJSON, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var postsList []typedefs.Posts

	// Parse the response.
	jsonErr := json.Unmarshal(responseJSON, &postsList)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	json.NewEncoder(w).Encode(postsList)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", ServePostList).Methods("GET")

	// Print it to the screen.
	http.ListenAndServe(":5000", router)



}