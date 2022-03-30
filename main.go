package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home - home page handler
func Home(w http.ResponseWriter, r *http.Request) {

}

// About - about page handler
func About(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Handlers registration
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// Start web-server
	_ = http.ListenAndServe(portNumber, nil)
}
