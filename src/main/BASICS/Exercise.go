package main

import (
	"fmt"
	"runtime"
)

func Sqrt(x float64) float64 {
	z := 1.0
	//miin := math.Abs(z*z - x)
	for count := 0; count < 10; count++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
