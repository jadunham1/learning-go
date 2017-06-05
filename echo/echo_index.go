package main

import (
    "fmt"
    "os"
)

func main() {
    for index, arg := range os.Args[1:] {
        fmt.Println(fmt.Sprintf("%d: %s", index+1, arg))
    }
}
