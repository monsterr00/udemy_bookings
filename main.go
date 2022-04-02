package main

import (
	"net/http"
)

const portNumber = ":8080"

func main() {
	// Handlers registration
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start web-server
	_ = http.ListenAndServe(portNumber, nil)
}
