package main

import "fmt"

type InterfaceI interface {
	Abs() float64
}

type Vertex struct {
	x float64
}

func (v Vertex) Abs() float64 {
	return v.x
}

/*
This function doesn't implement the interface
only methods can
*/
func Abs(v Vertex) float64 {
	return v.x
}

func basicInterface() {

}

func main() {
	var a InterfaceI
	a = Vertex{2}
	fmt.Println(a.Abs())
}
