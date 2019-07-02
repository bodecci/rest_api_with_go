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

// GET all Books
func getBooks(w http.ResponseWriter, r *http.Request) { //similar to (req, res)

}

// to GET single Book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// to create a new Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// to update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// to delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
