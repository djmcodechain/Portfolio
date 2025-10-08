package main

import (
	"danieljmanning/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.IndexHandler)

	// Serve Static Files
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
