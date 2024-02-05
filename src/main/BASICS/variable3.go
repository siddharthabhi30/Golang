package main

import (
	"fmt"
	"math"
)

func main() {
	b := 2.0
	fmt.Printf("Type: %T Value: %v\n", b, b)
	c := fmt.Sprint(math.Sqrt(b))

	fmt.Printf("Type: %T Value: %v\n", c, c)

}
