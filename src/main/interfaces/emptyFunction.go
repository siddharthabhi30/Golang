package main

type InterfaceI4 interface {
	Abs() float64
}

func main() {
	var a InterfaceI4
	a.Abs()
}
