package main

import (
  "log"
  "net/http"
  "os"
  "time"
  "io"
  "fmt"
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

func init() {
  trace = os.Stdout
  logger = log.New(trace, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
  logger.Println("init logger success")
}

func main() {
  http.HandleFunc("/", handler)
  logger.Println("Listening...")
  err := http.ListenAndServe(getPort(), nil)
  if err != nil {
    logger.Fatalln("ListenAndServe", err)
    return
  }
}

func getPort() string {
  var port = os.Getenv("PORT")
  if port == "" {
    port = "2424"
    logger.Println("No PORT environment variable detected, defaulting to", port)
  }
  return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
  logger.Println("Handler function called")
  fmt.Fprintf(w, "Hello. The time is now %s!", time.Now().Format("3:04:05"))
}
