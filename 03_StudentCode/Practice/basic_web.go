package main

import (
  "fmt"
  "net/http"
  "log"
  "io"
  "os"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
var (
  logger *log.Logger
  trace io.Writer
)

func handler(w http.ResponseWriter, r *http.Request) {
  logger.Println("handler function called")
  fmt.Fprint(w, "Hello World!")
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
  logger.Println("handler two function called")
  fmt.Fprint(w, "Hello Again!")
}

func handlerAgain(w http.ResponseWriter, r *http.Request) {
  logger.Println("handler golang called")
  fmt.Fprint(w, "Hello Golang")
}

//automatic call before main function
func init() {
  trace = os.Stdout
  logger = log.New(trace, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
  logger.Println("init function called")
  http.HandleFunc("/", handler)
  http.HandleFunc("/hello", handlerTwo)
  http.HandleFunc("/golang", handlerAgain)
}

func main() {
  logger.Println("main function called")
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatalln("ListenAndServe:", err)
  }
}
