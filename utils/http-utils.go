package utils

import (
	"io/ioutil"
	"net/http"
)

var BaseUrl string = "https://www.hotelflow.com.br/wp-json/wp/v2/"

func MakeRouteConfigs(w *http.ResponseWriter) {
	SetContentType(w)
	EnableCors(w)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func SetContentType(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func MakeGenericRequest(w http.ResponseWriter, r *http.Request, endpointSuffix string) []byte {
	MakeRouteConfigs(&w)
	// Endpoint to fetch the content from.
	endpoint := BaseUrl + endpointSuffix

	// Makes a GET request to the provided URL.
	response, err := http.Get(endpoint)
	LogError(err)

	// Read the response.
	responseJSON, readErr := ioutil.ReadAll(response.Body)
	LogError(readErr)

	// Return the response.
	return responseJSON
}