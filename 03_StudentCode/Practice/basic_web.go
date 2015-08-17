package main

import (
  "fmt"
  "net/http"
  "log"
  "io"
  "os"
)

var (
  logger *log.Logger
  trace io.Writer
)

func handler(w http.ResponseWriter, r *http.Request) {
  logger.Println("handler function called")
  fmt.Fprint(w, "Hello World!")
}

//automatic call before main function
func init() {
  trace = os.Stdout
  logger = log.New(trace, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
  logger.Println("init function called")
  http.HandleFunc("/", handler)
}

func main() {
  logger.Println("main function called")
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatalln("ListenAndServe:", err)
  }
}
