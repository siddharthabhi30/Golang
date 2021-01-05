package main

import (
	"fmt"
	"math/cmplx"
)

func add(x int, y int) int {
	return x + 23
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var c, python, java bool

func main() {
	var i int
	i = add(1, 2)
	// this is wrong--> u = add(1, 2)
	fmt.Println(i, c, python, java, i)

	var my = "sid"
	fmt.Print(my, MaxInt)
	fmt.Println()

	//var i2 int = 42
	//both ways are posisble for explicit conversion
	//var f float64 = float64(i)

	//f2 := float64(i2)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	v := 424.0 // change me!

	fmt.Printf("v is of type %T\n  val is %v", v, v)

	fmt.Println()
	const truth = true
	fmt.Printf("truth is of type %T\n  val is %v", truth, truth)

	fmt.Println(Big)
}
