package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var diff = float64(1000000000000)
	var valSelected float64

	for diff > 0.00000000000001 {
		z -= (z*z - x) / (2 * z)
		valSelected = z
		diff = min(diff, math.Abs(valSelected*valSelected-x))
	}
	return valSelected
}

func main() {
	fmt.Println(Sqrt(4))
}
