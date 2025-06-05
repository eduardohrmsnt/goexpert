package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request finished")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request completed")
		w.Write([]byte("Hello, World!"))
	case <-ctx.Done():
		log.Println("Request cancelled")
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
