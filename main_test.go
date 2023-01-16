package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
	got := "helo"
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}