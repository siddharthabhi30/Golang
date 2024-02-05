package main

import "fmt"

type InterfaceI3 interface {
	Abs() float64
}

type Vertex3 struct {
	x float64
}

func (v Vertex3) Abs() float64 {
	return v.x
}

func receiveConcreteAsParam(a Vertex3) {
	fmt.Printf(" value is %v\n", a.x)
}

/*
We are now type casting a variable
*/
func typeCasting() {
	var interfaceOnelevelVar InterfaceI3
	f, ok := interfaceOnelevelVar.(InterfaceI3)
	fmt.Println(f, ok)

	interfaceOnelevelVar = Vertex3{2}
	f2, ok2 := interfaceOnelevelVar.(InterfaceI3)
	fmt.Println(f2, ok2)

	interfaceOnelevelVar2 := interfaceOnelevelVar.(Vertex3)

	receiveConcreteAsParam(interfaceOnelevelVar2)
}

func main() {
	typeCasting()
}
