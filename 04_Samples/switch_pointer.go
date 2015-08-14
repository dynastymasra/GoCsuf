package main

import "fmt"

func main() {
  message := "Hello World Go Pointer!"

  var greeting *string = &message
  var greeting2 string = *greeting
  var greeting3 *string = &greeting2

  fmt.Println("message", message)
  fmt.Println("greeting", greeting)
  fmt.Println("greeting2", greeting2)
  fmt.Println("greeting3", *greeting3)  
}
