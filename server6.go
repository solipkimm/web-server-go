package main

import (
	"fmt"
	"net/http"
)

func WebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	mux.HandleFunc("/solip", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Solip!")
	})
	return mux
}

func main() {
	http.ListenAndServe(":3000", WebHandler())
}
