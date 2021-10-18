package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandleReq(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("test")
	w.Write([]byte(val))
}

func CheckResponse(r *http.Response) bool {
	return r.StatusCode == 200
}

func TestServer(t *testing.T) {
	http.HandleFunc("/", HandleReq)

	go http.ListenAndServe(":8080", nil)

	resp, err := http.Get("http://localhost:8080/?test=wowah")
	if err != nil {
		t.Fatalf("Send request error: %v", err)
	}

	if !CheckResponse(resp) {
		t.Fatalf("Response check failed")
	}
}

func TestServerTwo(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/?test=wowah", nil)
	w := httptest.NewRecorder()
	HandleReq(w, req)

	resp := w.Result()

	if !CheckResponse(resp) {
		t.Fatalf("Response check failed")
	}
}
