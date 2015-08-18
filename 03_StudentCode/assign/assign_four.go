package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
func main() {
  for i := 1; i < 6; i++ {
    fmt.Println(i)
  }

  fmt.Println("With Break")
  for i := 1; ; i++ {
    if i > 5 {
      break
    }
    fmt.Println(i)
  }

  fmt.Println("Even number")
  for i := 1; i < 6; i++ {
    if i % 2 == 1 {
      continue
    }
    fmt.Println(i)
  }

  nums := []int{1, 2, 3, 4, 5}
  fmt.Println("Slice")
  for _, item := range nums {
    fmt.Println(item)
  }

  myMap := map[string]string {
    "Kira" : "Yamato",
    "Athrun" : "Zala",
    "Mula" : "Flaga",
  }

  for index, item := range myMap {
    fmt.Println(index, item)
  }
}
