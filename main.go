package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// REST api should respond to the following actions:
//   create a todo, view a todo, modify a todo, list all todos, delete all todos

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	router.GET("/", createHandler(homeHandler))
	router.GET("/todos/:id", createHandler(showHandler))
	router.GET("/todos", createHandler(indexHandler))

	fmt.Printf("starting server on port %s\n", port)
	http.ListenAndServe(":"+port, router)
}

func homeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Todo Home")
}

func showHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET /:id") // fetch todo by id.
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("GET /") // fetch all todos.
	json.NewEncoder(w).Encode(mockApp.mockItem())
}
