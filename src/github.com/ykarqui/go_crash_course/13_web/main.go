package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About world")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server is starting")
	http.ListenAndServe(":3000", nil)
}
