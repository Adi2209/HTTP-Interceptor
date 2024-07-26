package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api1", apiHandler1)
	mux.HandleFunc("/api2", apiHandler2)

	compressionMiddleware := compressionInterceptor(mux)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", compressionMiddleware); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func apiHandler1(w http.ResponseWriter, r *http.Request) {
	response := "Hello from API 1!"
	w.Write([]byte(response))
}

func apiHandler2(w http.ResponseWriter, r *http.Request) {
	response := "Hello from API 2!"
	w.Write([]byte(response))
}
