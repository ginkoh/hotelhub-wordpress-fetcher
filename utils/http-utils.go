package utils

import "net/http"

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