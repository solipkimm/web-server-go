package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// get query
	values := r.URL.Query()
	// get a key value
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	// convert "id" value(string) to int
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! Your id is %d.", name, id)
}

func main() {
	// Set web handler
	http.HandleFunc("/", indexHandler)
	// Set web handler, port 3000, nil -> DefaultServeMux
	http.ListenAndServe(":3000", nil)
}
