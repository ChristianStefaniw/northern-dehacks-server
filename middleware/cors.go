package middleware

import (
	"net/http"
)

func CorsMiddleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "https://northerndehacks.ca")

	w.Header().Add("Access-Control-Allow-Headers", "*")
}
