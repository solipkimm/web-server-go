package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Set web handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	// Set web handler, port 3000, nil -> DefaultServeMux
	http.ListenAndServe(":3000", nil)
}
