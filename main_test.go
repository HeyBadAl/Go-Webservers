package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloWorldHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()

    // Use the updated exported handler function
    http.HandlerFunc(helloWorldHandler).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
    }

    expected := "Hello, World!"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expected)
    }
}

