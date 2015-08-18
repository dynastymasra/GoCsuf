package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
const(
  A = iota
  B = iota
  C = iota
  D = iota
  E = iota
  F = iota
  G = iota
)

type Message struct {
  author, body string
}

func show(message Message) {
  fmt.Println(message.author, "says:", message.body)
}

func multi(nums ...int)(int, int, int) {
  size := len(nums)
  return nums[size-3], nums[size-2], nums[size-1]
}

func main() {
  messageOne := Message{"Kira", "Freedom"}
  messageTwo := Message{"Athrun", "Justice"}
  messageThree := Message{"Shin", "Destiny"}

  _, secToLast, last := multi(A, B, C, D, E, F, G)
  thirdToLast, _, _ := multi(A, B, C, D, E, F, G)
  fmt.Println("thirdToLast", thirdToLast)
  fmt.Println("secToLast", secToLast)
  fmt.Println("last", last)

  show(messageOne)
  show(messageTwo)
  show(messageThree)
}
