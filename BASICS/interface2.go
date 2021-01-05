package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type I2 interface {
	M2()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	t.S = "69"
	fmt.Println(t.S)
}

func (t *T) M2() {

	t.S = "69"
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// my := T{"hello2"}

	// var i I = T{"hello"}

	// var i2 I2 = &my
	// i2.M2()
	// i.M()
	// fmt.Println(i)
	// fmt.Println(my)

	//main2()
	emptyInterface()
}
func main2() {
	//this i needsw some val ...else how to know which concrete method we arew implementing
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	//here we cant write var t T;
	var t *T
	i = t
	i.M()

	var tt T
	fmt.Println(tt)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
	var iii interface{}
	describe(iii)

	// iii = 42
	// describe(iii)

	// iii = "hello"
	// describe(iii)

}
