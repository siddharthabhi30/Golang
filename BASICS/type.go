package main

import "fmt"

func one() int {
	return 2
}
func two() int {
	return 3
}

type sttr string

func do(i interface{}) {

	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i interface{} = "hello"
	//t := i.(T) Here T refers to concrete type held by the interface
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	fmt.Println(f)

	ff := one()

	ff = two()
	fmt.Println(ff)
	do(21)
	do("hello")
	do(true)

	aa := 23

	// check := i.(type)
	// fmt.Println("check ", check) error

}
