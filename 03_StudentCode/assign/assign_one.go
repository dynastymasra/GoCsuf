package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
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
