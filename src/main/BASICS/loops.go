package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//score of v is in IF only and we are also excuting something
	// variable v here has limited scope. till else block only

	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {
		v2 := v
	} else {
		// we can now use v in all esle statement
		v3 := v
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	var name string = fmt.Sprint(23)
	fmt.Println(name)

}
