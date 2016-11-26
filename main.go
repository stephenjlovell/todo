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

	r := httprouter.New()
	// r.Handle("/", setHeader(HomeHandler))
	r.Handler("GET", "/", createHandler(todoHandler))

	fmt.Printf("starting server on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}

// func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	fmt.Fprintln(rw, "ToDo Home")
// }

func todoHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "ToDo Home")
	json.NewEncoder(w).Encode(mockApp.mockItem())
}
