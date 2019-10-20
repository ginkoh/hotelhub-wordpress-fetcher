package controllers

import (
	getters "../getters"
	typedefs "../typedefs"
	utils "../utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func ReturnParsedPostList(w http.ResponseWriter, r *http.Request) {
	// Initialize the struct variable to pass the JSON response as the content by reference.
	var postsList []typedefs.Posts

	// Parse the JSON response.
	jsonErr := json.Unmarshal(getters.GetPostList(w, r), &postsList)
	utils.LogError(jsonErr)

	for index, post := range postsList {
		var singleImage typedefs.FeaturedImage

		// Convert the featured media int to string.
		imageEndpoint := "media/" + strconv.Itoa(post.FeaturedMedia)

		// Get the post featured media.
		singleImgErr := json.Unmarshal(getters.GetSinglePostImage(w, r, imageEndpoint), &singleImage)
		utils.LogError(singleImgErr)

		// Assign the image link to the post.
		postsList[index].PostImage = singleImage.Guid.Rendered
	}

	// Return the encoded JSON to the router.
	json.NewEncoder(w).Encode(postsList)
}
