package main

import "fmt"

type InterfaceI2 interface {
	Abs() float64
}

type Vertex2 struct {
	x float64
}

func (v Vertex2) Abs() float64 {
	return v.x
}

func main() {
	checkOneLevelTypeInference()
}

/*
testing when a interface{} again points to interface
what happens then
*/
func basicInterfaceTypeCasting() {
	var interfaceVar interface{}
	var interfaceOnelevelVar InterfaceI2

	if interfaceVar == nil {
		fmt.Println("expected nil")
	}

	interfaceVar = interfaceOnelevelVar

	if interfaceVar == nil {
		fmt.Println("expected nil")
	}

	f, ok := interfaceVar.(InterfaceI2)
	fmt.Println(f, ok)

	interfaceOnelevelVar = Vertex2{2}
	// assigning again
	interfaceVar = interfaceOnelevelVar

	// This is a bit different.
	f2, ok2 := interfaceVar.(InterfaceI2)
	fmt.Println(f2, ok2)

	f3, ok3 := interfaceVar.(Vertex2)
	fmt.Println(f3, ok3)

}

/*
Until interface var points to a concrete type we
i.(type) won't give true
*/
func checkOneLevelTypeInference() {
	var interfaceOnelevelVar InterfaceI2
	f, ok := interfaceOnelevelVar.(InterfaceI2)
	fmt.Println(f, ok)

	interfaceOnelevelVar = Vertex2{2}
	f2, ok2 := interfaceOnelevelVar.(InterfaceI2)
	fmt.Println(f2, ok2)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
