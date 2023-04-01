package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's Go!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Println("Starting app on port 3000")
	// ListenAndServe always return non-nil err
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
