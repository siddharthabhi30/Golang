package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func testByte() {
	s := "hello"
	describe(s[0])
}

func main() {
	testByte()
}
