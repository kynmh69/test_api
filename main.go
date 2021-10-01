package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"test_api/model"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Bookのデータを保持するスライスの作成
var books []model.Book

const CantConnectDbMsg = "Can't connect DB.\n"

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&model.Book{})
	if err != nil {
		return
	}
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe in production
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book model.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("localhost:8083"))
	if err != nil {
		log.Fatalf(CantConnectDbMsg)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf(CantConnectDbMsg)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		return
	}

	// ルーターのイニシャライズ
	r := mux.NewRouter()

	// モックデータの作成
	books = append(books, model.Book{ID: "1", Title: "Book one", Author: &model.Author{FirstName: "Philip", LastName: "Williams"}})
	books = append(books, model.Book{ID: "2", Title: "Book Two", Author: &model.Author{FirstName: "John", LastName: "Johnson"}})

	// ルート(エンドポイント)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
