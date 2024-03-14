package handlers

import (
	"net/http"
)

// indexHandler is a simple HTTP handler function which writes a response.
// Serve the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "/home/jb/pixl/templates/index.html"
	http.ServeFile(w, r, filePath)
}
