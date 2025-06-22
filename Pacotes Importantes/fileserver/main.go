package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Static file server"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
