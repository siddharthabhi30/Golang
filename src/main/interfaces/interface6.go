package main

import "fmt"

type InterfaceI5 interface {
	Abs() float64
}

type Vertex5 struct {
	x float64
}

func (v Vertex5) Abs() float64 {
	return v.x
}

/*
We are now type casting a variable
*/
func typeCastingInterface() {
	var interfaceOnelevelVar InterfaceI5
	f, ok := interfaceOnelevelVar.(InterfaceI5)
	fmt.Println(f, ok)

	interfaceOnelevelVar = Vertex5{2}
	f2, ok2 := interfaceOnelevelVar.(InterfaceI5)
	fmt.Println(f2, ok2)

	//interfaceOnelevelVar2 := interfaceOnelevelVar.(Vertex5)

	var unknownVar interface{}
	unknownVar = interfaceOnelevelVar

	//interfaceOnelevelVar2 := interfaceOnelevelVar.(Vertex5)
	switch v := unknownVar.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Vertex5:
		fmt.Printf("this is %v interface 5 \n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

func main() {
	typeCastingInterface()
}
