package main

import (
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	got := os.Getenv("BIER")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}