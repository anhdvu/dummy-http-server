package main

import "net/http"

func enforceJSONHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		v, ok := headers["Content-Type"]
		if !ok || v[0] != "application/json" {
			http.Error(w, "request must have json content-type header.", http.StatusBadRequest)
			return
		}

		h.ServeHTTP(w, r)
	})
}
