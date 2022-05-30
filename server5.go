package main

import (
	"net/http"
)

func main() {
	// http://localhost:3000/static/gopher.jpeg
	// need to remove /static/ using StripPrefix
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
