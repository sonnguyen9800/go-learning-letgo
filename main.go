package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World! 1 1 1"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println(("Starting server on port 8080"))
	error := http.ListenAndServe(":8080", mux)
	log.Fatal(error)
}
