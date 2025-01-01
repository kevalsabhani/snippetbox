package main

import (
	"log"
	"net/http"
)

// different ways to run the server
// 1. go run main.go
// 2. go run .
// 3. go github.com/kevalsabhani/snippetbox

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Println("Starting server on :4000")

	log.Fatal(http.ListenAndServe(":4000", mux))
}
