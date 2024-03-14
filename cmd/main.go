// Set up and start web server
package main

import (
	"28jb11/pixl/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("/home/jb/pixl/public/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("/home/jb/pixl/public/js"))))

	// Set up routes and handlers
	http.HandleFunc("/pixels", handlers.PixelGridHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	// Start web server
	fmt.Println("Server is running at localhost:42069")
	if err := http.ListenAndServe(":42069", nil); err != nil {
		panic(err)
	}

}
