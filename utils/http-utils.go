package utils

import (
	typedefs "../typedefs"
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

func MakeAsyncRequests(urls []string) <-chan *typedefs.HttpResponse {
	// Create a channel (kind of a communicator between the go routines).
	channel := make(chan *typedefs.HttpResponse, len(urls))

	// Iterate over each of the given urls.
	for _, url := range urls {
		// Create a go routine (something like a thread) to make the requests simultaneously.
		go func(url string) {
			// Make the request.
			response, err := http.Get(url)

			// Read the response body.
			responseJSON, readErr := ioutil.ReadAll(response.Body)
			LogError(readErr)

			// Close the response body.
			response.Body.Close()

			// Send the response structure to the current channel.
			channel <- &typedefs.HttpResponse{url, responseJSON, err}
		}(url)
	}

	// Return the channel.
	return channel
}