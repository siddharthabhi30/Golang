package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func basicNil() {
	var i I
	if i == nil {
		// This would not be printed after (i = t)
		println("we found interace var to be nil")
	}
	var t *T
	i = t

	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func main() {
	basicNil()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
