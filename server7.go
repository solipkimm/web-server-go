package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

// WebHandler creates a handler instance
func WebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"Solip", 20, 85}

	// convert Student object to JSON-[]byte(an unsigned 8-bit integer)
	data, _ := json.Marshal(student)

	// set the type to JSON
	w.Header().Add("content-type", "application/json")

	// check the status
	w.WriteHeader(http.StatusOK)

	// send the result
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3000", WebHandler())
}
