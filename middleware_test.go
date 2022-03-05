package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_enforceJSONHeader(t *testing.T) {
	t.Run("No Content-Type header", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rq, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`"name":"Dace"`)))
		if err != nil {
			t.Fatal(err)
		}

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		enforceJSONHeader(next).ServeHTTP(rr, rq)

		rs := rr.Result()
		assertStatusCode(t, rs, http.StatusBadRequest)
	})

	t.Run("Wrong Content-Type header", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rq, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`"name":"Dace"`)))
		if err != nil {
			t.Fatal(err)
		}
		rq.Header.Set("Content-Type", "text/xml")

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		enforceJSONHeader(next).ServeHTTP(rr, rq)

		rs := rr.Result()
		assertStatusCode(t, rs, http.StatusBadRequest)
	})

	t.Run("Correct Content-Type header", func(t *testing.T) {
		rr := httptest.NewRecorder()

		rq, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`"name":"Dace"`)))
		if err != nil {
			t.Fatal(err)
		}
		rq.Header.Set("Content-Type", "application/json")

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		enforceJSONHeader(next).ServeHTTP(rr, rq)

		rs := rr.Result()
		assertStatusCode(t, rs, http.StatusOK)
	})
}
