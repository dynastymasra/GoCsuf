package math

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Mobile and Backend Developer
 */
func Average(intSliceParam []int) int {
  total := 0
  count := 3

  for _, x := range intSliceParam {
    total += x
  }
  return total / count
}

func AverageCorrect(intSliceParam []int) int {
  total := 0
  count := 0

  for _, x := range intSliceParam {
    total += x
    count++
  }
  return total / count
}
