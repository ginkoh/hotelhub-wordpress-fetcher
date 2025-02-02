package controllers

import (
	getters "../getters"
	typedefs "../typedefs"
	utils "../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
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
	for _ = range imagesEndpoints {
		// Create a image type (which has only one field) to hold the featured media link
		var singleImage typedefs.FeaturedImage

		// Receive the results response current channel.
		result := <-results

		// Parse the JSON adding each response to the singleImage structure.
		singleImgErr := json.Unmarshal(result.Response, &singleImage)
		utils.LogError(singleImgErr)

		// Initializes the found image flag.
		found := false

		// Iterate over the posts list to search for the current post image.
		for ind, post := range postsList {
			// Checks if the post id is the same as the image's post id.
			// While the correct image is not matched, continues the loop.
			if !found && post.FeaturedMedia == singleImage.ID {
				// Assign each post image to her post.

				if !strings.HasPrefix(singleImage.Guid.Rendered, "http") {
					postsList[ind].PostImage = utils.BaseSiteUrl + singleImage.Guid.Rendered
				} else {
					postsList[ind].PostImage = singleImage.Guid.Rendered
				}

				// The image was found, no need for new search in the current loop.
				found = true
			}
		}
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
