package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's Go!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing a snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Starting app on port 3000")
	// ListenAndServe always return non-nil err
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
