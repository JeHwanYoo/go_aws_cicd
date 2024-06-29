package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World %s", r.URL.Path[1:])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
