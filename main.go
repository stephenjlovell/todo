package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

// REST api should respond to the following actions:
//   create a todo, view a todo, modify a todo, list all todos, delete all todos

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("starting server on port %s\n", port)

	mux := http.NewServeMux()
	registerHandlers(mux)

	handler := cors.Default().Handler(mux)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		panic(err)
	}
}

func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", setHeaders(indexHandler))
	mux.HandleFunc("/:id", setHeaders(showHandler))
}

func setHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /:id") // fetch all todos.
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /:id") // fetch todo by id.
}
