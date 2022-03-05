package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_testHandler(t *testing.T) {
	cases := []struct {
		name     string
		method   string
		path     string
		body     []byte
		expected int
	}{
		{
			name:     "GET - Wrong path",
			method:   http.MethodGet,
			path:     "/test",
			body:     nil,
			expected: http.StatusNotFound,
		},
		{
			name:     "Get - Correct path",
			method:   http.MethodGet,
			path:     "/",
			body:     nil,
			expected: http.StatusMethodNotAllowed,
		},
		{
			name:     "POST - Wrong path",
			method:   http.MethodPost,
			path:     "/test",
			body:     []byte(`{"name":"Dace"}`),
			expected: http.StatusNotFound,
		},
		{
			name:     "POST - Correct path",
			method:   http.MethodPost,
			path:     "/",
			body:     []byte(`{"name":"Dace"}`),
			expected: http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var request *http.Request

			if c.body != nil {
				request, _ = http.NewRequest(c.method, c.path, bytes.NewBuffer(c.body))
			} else {
				request, _ = http.NewRequest(c.method, c.path, nil)
			}

			response := httptest.NewRecorder()

			testHandler().ServeHTTP(response, request)
			rs := response.Result()

			if rs.StatusCode != c.expected {
				t.Errorf("want %d; got %d", c.expected, rs.StatusCode)
			}
		})
	}

	t.Run("Response with application/json header", func(t *testing.T) {
		var request *http.Request

		request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"name":"Dace"}`)))

		response := httptest.NewRecorder()

		testHandler().ServeHTTP(response, request)
		rs := response.Result()

		if cth := rs.Header.Get("Content-Type"); cth != "application/json" {
			t.Errorf("want %s; got %s", "application/json", cth)
		}

		rsb, _ := io.ReadAll(rs.Body)
		expectedBody := `{"name":"Dace","response":"approved"}`

		if string(rsb) != expectedBody {
			t.Errorf("want %s; got %s", expectedBody, string(rsb))
		}
	})

}

func assertStatusCode(t *testing.T, rs *http.Response, expected int) {
	t.Helper()
	if rs.StatusCode != expected {
		t.Errorf("want %d; got %d", expected, rs.StatusCode)
	}
}
