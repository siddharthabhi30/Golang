package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	var diff = float64(1000000000000)
	var valSelected float64

	for diff > 0.00000000000001 {
		z -= (z*z - x) / (2 * z)
		valSelected = z
		diff = min(diff, math.Abs(valSelected*valSelected-x))
	}
	return valSelected, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
