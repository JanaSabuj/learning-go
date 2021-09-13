package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define routes
	http.HandleFunc("/greet", greet)

	// starting the server
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
