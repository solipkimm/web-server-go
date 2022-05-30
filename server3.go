package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a ServeMux instance
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	mux.HandleFunc("/solip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Solip!")
	})
	http.ListenAndServe(":3000", mux)
}
