package math

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
func Sum(intSliceParam []int) int {
  total := 0

  for _, x := range intSliceParam {
    total += x
  }
  return total
}
