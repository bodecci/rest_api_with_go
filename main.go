package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Sruct (Model)
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"id"`
}

// Author struct
type Author struct {
	FirstName string `json:"firtstname"`
	LastName  string `json:"lastname"`
}

func main() {
	// init router
	r := mux.NewRouter()

	// route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id", getBooks).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
