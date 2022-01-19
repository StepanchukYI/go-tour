package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error with number: %.f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := float64(1)
	i := 0
	for ; i < 10000; i++ {
		z = z - ((math.Pow(z, 2.0) - x) / 2 * z)
	}
	fmt.Println("Number of iterations:", i)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
