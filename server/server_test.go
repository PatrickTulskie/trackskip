package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		inputURL         string
		expectedLocation string
		expectedStatus   int
	}{
		{
			inputURL:         "/?q=https%3A%2F%2Fgoogle.com",
			expectedLocation: "https://google.com",
			expectedStatus:   http.StatusSeeOther,
		},
		// Add more tests cases as required
	}

	for _, tt := range tests {
		req, err := http.NewRequest("GET", tt.inputURL, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		// Use httptest to create a ResponseRecorder (which satisfies the http.ResponseWriter interface)
		// This way you can inspect the response afterwards
		rec := httptest.NewRecorder()

		handler(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		if res.StatusCode != tt.expectedStatus {
			t.Errorf("expected status %d; got %d", tt.expectedStatus, res.StatusCode)
		}

		location, ok := res.Header["Location"]
		if ok {
			if len(location) > 0 && location[0] != tt.expectedLocation {
				t.Errorf("expected location %s; got %s", tt.expectedLocation, location[0])
			}
		} else {
			t.Errorf("location header not set")
		}
	}
}
