package main

import (
  "errors"
  "fmt"
  "io"
  "net/http"
  "os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("got / request\n")
  io.WriteString(w, "This is a Website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
  http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
}
