package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		fmt.Fprintf(w, "Application is up")
	})

	// Start the HTTP server and listen on port 8080
	// http.ListenAndServe blocks until the program is terminated
	fmt.Println("Server listening on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
