package main_test

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Println("hello world2")
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}