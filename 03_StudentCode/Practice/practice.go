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

func(s str) testPrint(t str) {
  fmt.Println(s, t)
}

func printSlice(slice []int) {
  for _, current := range slice {
    fmt.Println(current)
  }
  fmt.Println()
}

func main()  {
  var myVar [6]int
  other := myVar[:5]
  for key, _ := range other {
    other[key] = key + 1
  }

  fmt.Println("Lenght:", len(myVar), "\nCapacity:", cap(other))
  printSlice(other)
  other = append(other, 6)
  other = append(other, 7, 8)
  printSlice(other)

  for i := range myVar {
    fmt.Println(myVar[i])
  }
  var myStr str = "Hello World!"
  myStr.testPrint("Now Learn Go")
}
