package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

var books = []string{"The Lord of the Rings", "The Hobbit", "The Silmarillion"}

type Book struct {
	Title string `json:"title"`
}

func main() {
	err := runServer()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server shutdown")
		} else {
			fmt.Println("server failed", err)
		}
	}
}

func runServer() error {
	httpRouter := chi.NewRouter()

	httpRouter.Route("/api/v1", func(r chi.Router) {
		r.Get("/books", getAllBooks)
		r.Post("/books", addBook)
		r.Delete("/books", deleteAllBooks)
	})

	server := &http.Server{Addr: "localhost:8888", Handler: httpRouter}
	return server.ListenAndServe()
}

func getAllBooks(w http.ResponseWriter, req *http.Request) {
	respBody, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func addBook(w http.ResponseWriter, req *http.Request) {
	var book Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books = append(books, book.Title)

	w.WriteHeader(http.StatusCreated)
}

func deleteAllBooks(w http.ResponseWriter, req *http.Request) {
	books = []string{}

	w.WriteHeader(http.StatusNoContent)
}
