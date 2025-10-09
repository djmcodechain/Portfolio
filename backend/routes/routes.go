package routes

import (
	"danieljmanning/backend/handlers"
	"net/http"
)

func Routes() {
	Mux := http.NewServeMux()

	Mux.HandleFunc("/", handlers.IndexHandler)

	http.ListenAndServe(":8080", Mux)
}
