package main

import "fmt"

type Name struct {
  firtName, lastName string
}

func main() {
  var message string = "Hello, my name is"
  var greeting *string = &message
  var me = Name{"Kira", "Yamato"}
  fmt.Println(message, me.firtName, me.lastName)

  *greeting = "You may call me"
  fmt.Println(message, me.firtName)
}
