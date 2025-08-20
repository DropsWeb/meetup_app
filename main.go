package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		hostname := GetLocalIP()
		fmt.Fprintf(w, "Application is up, host is: %v\n", hostname)
	})

	// Start the HTTP server and listen on port 8080
	// http.ListenAndServe blocks until the program is terminated
	fmt.Println("Server listening on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
