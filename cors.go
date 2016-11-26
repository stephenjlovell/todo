package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setHeader(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("access-control-allow-headers", "accept, content-type")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if r.Method != "OPTIONS" {
			handler(w, r, p)
		}
	}
}

func createHandler(handler httprouter.Handle) httprouter.Handle {
	return setHeader(handler)
}
