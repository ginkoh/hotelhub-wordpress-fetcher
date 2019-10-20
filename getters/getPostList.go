package getters

import (
	utils "../utils"
	"net/http"
)

func GetPostList(w http.ResponseWriter, r *http.Request) []byte {
	return utils.MakeGenericRequest(w, r, "posts")
}

func GetSinglePost(w http.ResponseWriter, r *http.Request, postId string) []byte {
	return utils.MakeGenericRequest(w, r, postId)
}