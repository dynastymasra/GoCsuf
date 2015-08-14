package main

import "fmt"

func main() {
  message := "Hello World!"
  fmt.Println("Message:", message)

  var greeting *string = &message
  fmt.Println("Greeting:", greeting)
  *greeting = "New One"

  fmt.Println("Message and Greeting:", message, *greeting)
}
