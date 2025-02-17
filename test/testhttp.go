package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, from Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil) // Run on 8080 instead of 80
}
