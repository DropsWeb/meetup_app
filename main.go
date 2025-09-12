package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		time.Sleep(1 * time.Second)
		result, err := readFromDatabase()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "%v\n", result)
		fmt.Printf("%v\n", result)
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
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func readFromDatabase() (string, error) {
	rand.Seed(time.Now().UnixNano())

	filename := fmt.Sprintf("%d.txt", rand.Intn(100000))

	f, err := os.Create(filename)

	if err != nil {
		return "", err
	}

	defer f.Close()

	val := fmt.Sprintf("Результат чтения из БД. IP ноды приложения: %s", GetLocalIP())

	if _, err := f.WriteString(val); err != nil {
		return "", err
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
