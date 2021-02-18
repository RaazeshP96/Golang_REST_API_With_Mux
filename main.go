package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Book struck

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init book slice
var books []Book

//controller
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewDecoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = item.ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}

	}
	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}

//book route
func main() {
	godotenv.Load()
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "444874", Title: "Book on go", Author: &Author{
		Firstname: "John", Lastname: "shrestha"}})
	books = append(books, Book{ID: "2", Isbn: "444874", Title: "Book on python", Author: &Author{
		Firstname: "Mc Author", Lastname: "Smith"}})

	r.HandleFunc("/api/books", getBooks).Methods(("GET"))
	r.HandleFunc("/api/book", createBook).Methods(("POST"))
	r.HandleFunc("/api/book/{id}", getBookByID).Methods(("GET"))
	r.HandleFunc("/api/book/{id}", updateBook).Methods(("PUT"))
	r.HandleFunc("/api/book/{id}", deleteBook).Methods(("DELETE"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
