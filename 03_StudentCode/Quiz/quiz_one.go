package main

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
type str string

func(s str) printOne(i int) {
  fmt.Println(s, "likes", i)
}

func main() {
  nums := map[str]int {
    "Kira Yamato" : 20, "Lacas Cylane" : 21, "Mula Flaga" : 22, "Mariu Ramias" : 23,
    "Athrun Zala" : 24, "Cagali Yula Atha" : 25,
  }

  for key, value := range nums {
    key.printOne(value)
  }
}
