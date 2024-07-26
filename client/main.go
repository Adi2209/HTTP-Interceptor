package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang/snappy"
)

func main() {
	url := "http://localhost:8080"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request: %v\n", err)
		os.Exit(1)
	}

	req.Header.Set("Accept-Encoding", "snappy")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var responseBody []byte
	if resp.Header.Get("Content-Encoding") == "snappy" {
		snappyReader := snappy.NewReader(resp.Body)
		responseBody, err = io.ReadAll(snappyReader)
		if err != nil {
			fmt.Printf("Failed to read snappy response: %v\n", err)
			os.Exit(1)
		}
	} else {
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed to read response: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Response: %s\n", responseBody)
}
