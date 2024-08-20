package main

import (
	"log"
	"net/http"
	"strconv"
)

const PORT = 9999

//#endregion

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/hello", hello)

	mux.HandleFunc("/snippet/view/", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println(("Starting server on port " + strconv.Itoa(PORT)))
	error := http.ListenAndServe(":"+strconv.Itoa(PORT), mux)
	log.Fatal(error)
}
