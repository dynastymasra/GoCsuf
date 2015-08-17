package main

import "fmt"

type Operator func(int, int) int

func operation(numOne, numTwo int, operator Operator) int {
  result := operator(numOne, numTwo)
  if result < 0 {
    fmt.Println("Negative")
  } else {
    fmt.Println("Positive")
  }
  return result
}

func creation() (Operator, Operator, Operator, Operator) {
  return func(a, b int) int {
    return a + b
  }, func(a, b int) int {
    return a - b
  }, func(a, b int) int {
    return a * b
  }, func(a, b int) int {
    return a / b
  }
}

func main() {
  add, subtract, multiply, divide := creation()
  a, b := 6, 3
  var result int
  chose := "multiply"
  if a < b {
    chose = "subtract"
  }

  switch chose {
  case "add":
    result = operation(a, b, add)
  case "subtract":
    result = operation(a, b, subtract)
  case "multiply":
    result = operation(a, b, multiply)
  case "divide":
    result = operation(a, b, divide)
  }

  fmt.Println("Result", result)
}
