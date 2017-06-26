package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        prefix := "http://"
        if !strings.HasPrefix(url, prefix) {
            url = prefix + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }
        if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
            fmt.Fprintf(os.Stderr, "copy problem: %v\n", err)
        }
    }
}
