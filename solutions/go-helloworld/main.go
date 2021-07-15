package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {su
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":6111", nil)
}
