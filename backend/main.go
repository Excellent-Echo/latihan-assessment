package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	port := "localhost:8000"

	fmt.Println("Starting web server at", port)
	http.ListenAndServe(port, nil)
}
