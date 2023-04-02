package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Starting app on port 4000")
	// ListenAndServe always return non-nil err
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
