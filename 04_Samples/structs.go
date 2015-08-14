package main

import "fmt"

type Identity struct {
  firstname string
  lastname string
  age int
}

func main() {
  var c = Identity{"Atran", "Zala", 20}
  var b = Identity{firstname:"Kira", lastname:"Yamato", age:20}
  var a = Identity{}
  a.firstname = "Mula"
  a.lastname = "Flaga"
  a.age = 25

  fmt.Printf("Firstname : %v, Lastname : %v, Age : %v \n", c.firstname, c.lastname, c.age)
  fmt.Printf("Firstname : %v, Lastname : %v, Age : %v \n", b.firstname, b.lastname, b.age)
  fmt.Printf("Firstname : %v, Lastname : %v, Age : %v \n", a.firstname, a.lastname, a.age)  
}
