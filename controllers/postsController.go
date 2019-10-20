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

	// Create a holder array to the endpoints.
	var imagesEndpoints []string

	// Iterate over the already plotted (with data from wordpress) posts list.
	// _ is ignored in all places.
	for _, post := range postsList {
		// Convert the featured media int to string to form a link.
		imageEndpoint := utils.BaseUrl + "media/" + strconv.Itoa(post.FeaturedMedia)
		// Append the current image endpoint to the endpoints list.
		imagesEndpoints = append(imagesEndpoints, imageEndpoint)
	}

	// Make the asynchronous requests for the media links.
	results := utils.MakeAsyncRequests(imagesEndpoints)

	// Iterate over the endpoints.
	for index := range imagesEndpoints {
		// Create a image type (which has only one field) to hold the featured media link
		var singleImage typedefs.FeaturedImage

		// Receive the results response current channel.
		result := <-results

		// Parse the JSON adding each response to the singleImage structure.
		singleImgErr := json.Unmarshal(result.Response, &singleImage)
		utils.LogError(singleImgErr)

		// Assign each post image to her post.
		postsList[index].PostImage = singleImage.Guid.Rendered
	}

	// Send the encoded response to the router.
	json.NewEncoder(w).Encode(postsList)
}

func ReturnSingleParsedPost(w http.ResponseWriter, r *http.Request) {
	// Create structures to hold the fetch data.
	var singlePost typedefs.Posts
	var singleImage typedefs.FeaturedImage

	// Get the post id param.
	params := mux.Vars(r)
	postId := params["postId"]

	// Parse the JSON.
	jsonErr := json.Unmarshal(getters.GetSinglePost(w, r, postId), &singlePost)
	utils.LogError(jsonErr)

	// Convert the featured media int to string.
	imageEndpoint := "media/" + strconv.Itoa(singlePost.FeaturedMedia)
	// Make the request for the post image.
	singleImgErr := json.Unmarshal(getters.GetSinglePostImage(w, r, imageEndpoint), &singleImage)
	utils.LogError(singleImgErr)

	// Assign the post image to the current single post.
	singlePost.PostImage = singleImage.Guid.Rendered

	// Send the encoded response to the router.
	json.NewEncoder(w).Encode(singlePost)
}