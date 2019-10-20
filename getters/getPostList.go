package getters

import (
	utils "../utils"
	"io/ioutil"
	"net/http"
)

func GetPostList(w http.ResponseWriter, r *http.Request) []byte {
	utils.MakeRouteConfigs(&w)

	// Endpoint to fetch the content from.
	endpoint := baseUrl + "posts"

	// Makes a GET request to the provided URL.
	response, err := http.Get(endpoint)
	utils.LogError(err)

	// Read the response.
	responseJSON, readErr := ioutil.ReadAll(response.Body)
	utils.LogError(readErr)

	// Return the response.
	return responseJSON
}