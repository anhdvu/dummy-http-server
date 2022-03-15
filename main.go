package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("I am a dummy HTTP server and I only want POST.")

	mux := http.NewServeMux()
	mux.Handle("/callback", enforceJSONHeader(testHandler()))

	srv := &http.Server{
		Addr:         ":8989",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("I am listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
