package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT = 9999

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello"))
}

// #region handle snippet
func viewSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.Error(w, "Invalid snippet ID", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Snippet: View %d", id)
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("X-Whatever", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}
	w.Write([]byte("Snippet: Create"))
}

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
