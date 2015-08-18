package main

import (
  "fmt"
  "github.com/dynastymasra/practice/slicelib"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
*/
type StrString string

type Name struct {
  fisrtname StrString
  lastname StrString
}

func(s StrString) initial() string {
  return string(s[0])
}

func(n Name) fullname() StrString {
  return n.fisrtname + " " + n.lastname
}

func printslice(s []string) {
  for _, item := range s {
    fmt.Println(item)
  }
}

func main() {
  gundam := slicelib.StrSlice{"Zeta", "Wings", "Seed", "Destiny", "00", "Unicorn"}
  gundamArray := [6]string{"Zeta", "Wings", "Seed", "Destiny", "00", "Unicorn"}
  same := true
  for index, item := range gundam {
    if item == gundamArray[index] {
      continue
    } else {
      same = false
      fmt.Println("Not equal")
    }
  }
  if same {
    fmt.Println("All slice equal")
  }
  fmt.Println("All available gundam")
  printslice(gundam)

  fmt.Println("First gundam")
  gundam.PrintFirst()

  secondHalf := gundam.SecondHalf()
  fmt.Println("Second Half gundam")
  printslice(secondHalf)

  axis := gundam[:2]
  axis = append(axis, gundam[2:4]...)
  fmt.Println("Gundam:")
  printslice(axis)

  ms := gundam[:]
  ms = append(ms[:4], ms[5:]...)
  fmt.Println("Mobile Suit")
  printslice(ms)

  name := Name{"Kira", "Yamato"}
  fmt.Println("Name is", name.fullname())
  fmt.Println("Initial is", name.fisrtname.initial(), name.lastname.initial())
}
