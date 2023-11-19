package main

import "fmt"

func test() (string, string) {
	return "a", "b"
}

func main() {
	// We can assign variable differently
	//var a string
	//var b string
	a, b := test()
	fmt.Print(a)
	fmt.Print(b)
}
