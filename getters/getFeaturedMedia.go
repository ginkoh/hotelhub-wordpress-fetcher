package getters

import (
	"net/http"
	utils "../utils"
)

func GetSinglePostImage(w http.ResponseWriter, r *http.Request, suffix string) []byte {
	return utils.MakeGenericRequest(w, r, suffix)
}