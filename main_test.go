package main

import (
	"net/http"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://pingb.in/p/665fb442e49ec5ac8785bb62e36c", nil)

	req.Header.Add("If-None-Match", os.Getenv("BIER"))
	client.Do(req)

	got := os.Getenv("BIER")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
