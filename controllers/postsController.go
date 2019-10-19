package controllers

import (
	getters "../getters"
	typedefs "../typedefs"
	utils "../utils"
	"encoding/json"
	"net/http"
)

func ReturnParsedPostList(w http.ResponseWriter, r *http.Request) {
	// Initialize the struct variable to pass the JSON response as the content by reference.
	var postsList []typedefs.Posts

	// Parse the JSON response.
	jsonErr := json.Unmarshal(getters.GetPostList(w, r), &postsList)
	utils.LogError(jsonErr)

	// Encode the JSON.
	json.NewEncoder(w).Encode(postsList)
}
