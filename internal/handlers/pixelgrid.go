package handlers

import (
	"28jb11/pixl/internal/templates"
	"net/http"
)

// PixelGridHandler serves the pixel grid HTML page
func PixelGridHandler(w http.ResponseWriter, r *http.Request) {

	// w.Write([]byte("pixels"))
	// Correctly point to the location of pixelgrid.html
	// filePath := filepath.Join("templates", "pixelgrid.html")
	// filePath := "/home/jb/pixl/templates/pixelgrid.html"
	// http.ServeFile(w, r, filePath)
	//
	templates.PixelGridPage(w, r)
}
