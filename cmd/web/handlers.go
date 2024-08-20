package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/pages/error.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "error", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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
