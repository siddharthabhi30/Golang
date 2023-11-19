package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Method and functions do copy in golang language
*/
func main() {
	funcReferenceTest()
}

func basicFunc() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func (v Vertex) Abs2() Vertex {
	v.X = 99
	return v
}

func Abs2PureFunc(v Vertex) Vertex {
	v.X = 999
	return v
}

func funcReferenceTest() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs2())
	fmt.Println(v)

	fmt.Println(Abs2PureFunc(v))
	fmt.Println(v)

}
