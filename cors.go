package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setHeader(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Processing request...")
		if r.Method == "OPTIONS" {
			fmt.Println("Preflight...")
			w.Header().Set("access-control-allow-headers", "X-Requested-With")
		} else {
			w.Header().Set("access-control-allow-origin", "*")
			w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("access-control-allow-headers", "accept, content-type")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		}
		handler(w, r, p)
	}
}

func createHandler(handler httprouter.Handle) httprouter.Handle {
	fmt.Println("Registering new handler...")
	return setHeader(handler)
}

// curl -H "Origin: http://www.example.com" \
//   -H "Access-Control-Request-Method: POST" \
//   -H "Access-Control-Request-Headers: X-Requested-With" \
//   -X OPTIONS --verbose \
//   http://localhost:8080
