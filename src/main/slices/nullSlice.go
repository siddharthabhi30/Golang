package main

import "fmt"

/*
Null slice is slice of length 1
*/
func main() {
	var a []int
	a = append(a, 99)
	fmt.Print(len(a))
}
