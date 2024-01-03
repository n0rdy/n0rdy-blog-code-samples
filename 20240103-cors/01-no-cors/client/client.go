package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := runServer()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("client server shutdown")
		} else {
			fmt.Println("client server failed", err)
		}
	}
}

func runServer() error {
	httpRouter := chi.NewRouter()

	httpRouter.Get("/", serveIndex)

	server := &http.Server{Addr: "localhost:3333", Handler: httpRouter}
	return server.ListenAndServe()
}

func serveIndex(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "20240103-cors/01-no-cors/client/index.html")
}
