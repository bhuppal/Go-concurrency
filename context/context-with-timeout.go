package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// set a http client timeout

	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	if err != nil {
		log.Fatal(err)
	}

	// creating a context with timeout and setting 100 millisecond
	ctx, cancel := context.WithTimeout(req.Context(), 1 * time.Millisecond)
	defer cancel()

    req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// Close the response body on the return
	defer resp.Body.Close()

	// Write the response to Stdout
	io.Copy(os.Stdout, resp.Body)
}
