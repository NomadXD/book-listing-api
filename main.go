package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Book represents a book entity.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Define HTTP endpoints
	http.HandleFunc("/books", listBooks)
	http.HandleFunc("/books/add", addBook)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/delete", deleteBook)

	// Start the HTTP server
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.SetKeepAlivesEnabled(false)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// listBooks handles the GET request to list all books.
func listBooks(w http.ResponseWriter, r *http.Request) {
	// Hardcoded list of books
	books := []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1"},
		{ID: 2, Title: "Book 2", Author: "Author 2"},
	}

	// Convert books to JSON and send as response
	jsonBytes, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// addBook handles the POST request to add a new book.
func addBook(w http.ResponseWriter, r *http.Request) {
	// Hardcoded response for adding a book
	newBook := Book{ID: 3, Title: "New Book", Author: "New Author"}

	// Respond with the newly added book
	jsonBytes, err := json.Marshal(newBook)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
}

// updateBook handles the PUT request to update a book.
func updateBook(w http.ResponseWriter, r *http.Request) {
	// Hardcoded response for updating a book
	updatedBook := Book{ID: 1, Title: "Updated Book", Author: "Updated Author"}

	// Respond with the updated book
	jsonBytes, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// deleteBook handles the DELETE request to delete a book by ID.
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// Hardcoded response for deleting a book
	w.WriteHeader(http.StatusNoContent)
}
