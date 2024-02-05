package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
we can't do this with primitive type
*/
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	referencePass()
}

func basicMethod() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

}

func referencePass() {
	v := Vertex{3, 4}
	(&v).Scale(10)
	fmt.Println(v.Abs())
}
