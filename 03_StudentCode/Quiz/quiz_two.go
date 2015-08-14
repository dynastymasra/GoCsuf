package main

import "fmt"

func main() {
  gundam := []string{"Wings", "Seed", "Destiny", "Unicorn", "00", "Stargazer", "Zeta"}
  fmt.Println("List of Gundam:")
  for _, list := range gundam {
    fmt.Println(list)
  }

  watch := gundam[:4]
  fmt.Println("Gundam in Show:")
  for _, list := range watch {
    fmt.Println(list)
  }

  fmt.Println("[:]", gundam[:])
  fmt.Println("[:2]", gundam[:2])
  fmt.Println("[4:]", gundam[4:])
}
