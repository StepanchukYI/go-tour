package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := float64(1)
  i := 0
  for ; i < 10000; i++ {
    z = z - ((math.Pow(z, 2.0) - x) / 2 * z)
  }
  fmt.Println("Number of iterations:", i)
  return z
}
func main() {
  fmt.Println(Sqrt(4))
  fmt.Println(Sqrt(6))
  fmt.Println(Sqrt(3))
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(4))
}