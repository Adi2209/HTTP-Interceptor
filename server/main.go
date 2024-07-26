package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang/snappy"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := "Hello, this is a Snappy compressed response!"

	if r.Header.Get("Accept-Encoding") == "snappy" {
		var buf bytes.Buffer
		snappyWriter := snappy.NewBufferedWriter(&buf)
		if _, err := snappyWriter.Write([]byte(response)); err != nil {
			http.Error(w, "Failed to write snappy response", http.StatusInternalServerError)
			return
		}
		snappyWriter.Close()

		w.Header().Set("Content-Encoding", "snappy")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, &buf)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}
}
