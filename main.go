package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write "Hello, World!" to the response writer
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		cores := runtime.NumCPU()
		runtime.GOMAXPROCS(cores)

		var wg sync.WaitGroup
		wg.Add(cores)

		for c := 0; c < cores; c++ {
			go func() {
				defer wg.Done()
				for i := 0; i < 10; i++ {
					res := mathHeavy(10_000_000)
					// Используем результат, чтобы компилятор не выкинул
					if res == 0 {
						fmt.Println("Impossible")
					}
				}
			}()
		}

		wg.Wait()

		result, err := readFromDatabase()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			fmt.Println(err.Error())
		}

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "%v\n", result)
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

func mathHeavy(iter int) float64 {
	result := 0.0
	for i := 1; i <= iter; i++ {
		result += math.Sqrt(float64(i)) * math.Sin(float64(i))
	}
	return result
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
