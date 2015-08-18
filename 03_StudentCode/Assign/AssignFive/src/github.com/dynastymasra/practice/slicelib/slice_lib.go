package slicelib

import "fmt"

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
type StrSlice []string

func(s StrSlice) PrintFirst() {
  fmt.Println(s[0])
}

func(s StrSlice) SecondHalf() StrSlice {
  return s[(len(s) / 2):]
}
