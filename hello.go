package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struck (Model)
type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Author Struck (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func findOne(w http.ResponseWriter, r *http.Request) {

}
func insert(w http.ResponseWriter, r *http.Request) {

}
func update(w http.ResponseWriter, r *http.Request) {

}
func delete(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{
		ID:     "1",
		Isbn:   "12887347",
		Title:  "Bukan Manusia Biasa",
		Author: &Author{FirstName: "Muhammad", LastName: "Al-Pandi"},
	})

	books = append(books, Book{
		ID:    "2",
		Isbn:  "3432562",
		Title: "Manusia Harimau",
	})

	r.HandleFunc("/api/books", findAll).Methods("GET")
	r.HandleFunc("/api/book/{id}", findOne).Methods("GET")
	r.HandleFunc("/api/book", insert).Methods("POST")
	r.HandleFunc("/api/book/{id}", update).Methods("PUT")
	r.HandleFunc("/api/book/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
