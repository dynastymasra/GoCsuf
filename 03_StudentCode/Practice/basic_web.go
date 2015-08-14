package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World!")
}

func init() {
  fmt.Println("init Function")
  http.HandleFunc("/", handler)
}

func main() {

}
