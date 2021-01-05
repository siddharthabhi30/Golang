package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	fmt.Println("hereee")
	//we are bracketing e with float coz conversion of e to string is not known ...so it will again call the same methid
	//initiating infinite loop
	return fmt.Sprintf("cannot Sqrt negative number %v ff", float64(e))
	//return "hhhh"
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	fmt.Println("hello")
	z := float64(2.)
	s := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s, nil
}

func main() {
	//fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
