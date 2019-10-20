package controllers

import (
	getters "../getters"
	typedefs "../typedefs"
	utils "../utils"
	"encoding/json"
	"github.com/gorilla/mux"
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
		singleImgErr := json.Unmarshal(getters.GetSinglePostImage(w, r, imageEndpoint), &singleImage)
		utils.LogError(singleImgErr)

		// Assign the image link to the post.
		postsList[index].PostImage = singleImage.Guid.Rendered
	}

	// Return the encoded JSON to the router.
	json.NewEncoder(w).Encode(postsList)
}

func ReturnSingleParsedPost(w http.ResponseWriter, r *http.Request) {
	var singlePost typedefs.Posts
	var singleImage typedefs.FeaturedImage

	params := mux.Vars(r)
	postId := params["postId"]

	jsonErr := json.Unmarshal(getters.GetSinglePost(w, r, postId), &singlePost)
	utils.LogError(jsonErr)

	// Convert the featured media int to string.
	imageEndpoint := "media/" + strconv.Itoa(singlePost.FeaturedMedia)
	singleImgErr := json.Unmarshal(getters.GetSinglePostImage(w, r, imageEndpoint), &singleImage)
	utils.LogError(singleImgErr)

	singlePost.PostImage = singleImage.Guid.Rendered

	json.NewEncoder(w).Encode(singlePost)
}