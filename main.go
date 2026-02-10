package main

import (
	"fmt"
	"net/http"
)

func blueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Blue page!")
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the Red page!")
}

func main() {
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)
	http.ListenAndServe(":8080", nil)
}
