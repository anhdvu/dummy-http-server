package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func testHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/callback" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "only POST method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		payload := make(map[string]interface{})

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("%+v\n", payload)
		payload["response"] = "approved"

		writeResponse(w, payload)
	})
}

func writeResponse(w http.ResponseWriter, pl interface{}) {
	rpBody, err := json.Marshal(pl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rpBody)
}
