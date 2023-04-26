package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/test", HandlePost1)
	log.Fatal(http.ListenAndServe(":8180", nil))

	time.Sleep(30 * time.Second)

}

func HandlePost1(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received: %s, %s\n ", r.Method, r.RequestURI)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")

}

/*
func processPutRequest(req *http.Request) []byte {
	// Get the chunked transfer encoding header.
	transferEncoding := req.Header.Get("Transfer-Encoding")

	// Check if the request is using chunked transfer encoding.
	if transferEncoding != "chunked" {
		return nil
	}

	// Create an array to store the chunks.
	chunks := make([]byte, 0)

	// Read the chunks from the request body.
	for {
		// Read the next chunk header.
		chunkHeader, err := req.Body.ReadBytes('\n')
		if err != nil {
			break
		}

		// Get the chunk length.
		chunkLength, err := strconv.Atoi(string(chunkHeader[:len(chunkHeader)-2]))
		if err != nil {
			break
		}

		// Read the chunk body.
		chunk, err := req.Body.ReadN(chunkLength)
		if err != nil {
			break
		}

		// Append the chunk to the array.
		chunks = append(chunks, chunk...)
	}

	// Return the array of chunks.
	return chunks
}
*/
