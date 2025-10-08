package main

import (
	"danieljmanning/backend/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)

	// Serve Static Files
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
