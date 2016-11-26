package main

import "net/http"

func setHeader(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("access-control-allow-headers", "accept, content-type")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if r.Method == "OPTIONS" {
			// handle preflight?
			return
		}
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func createHandler(handler http.HandlerFunc) http.Handler {
	return setHeader(handler)
}
