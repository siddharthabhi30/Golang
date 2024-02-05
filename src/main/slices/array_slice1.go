package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*
this exp shows that the
lower index cannot go left but right index can change

although, don't use this technique.
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[0:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:4]
	printSlice(s)

	s = s[0:4]
	printSlice(s)
	testNil()
}

/*
slices can be nil but not array
*/
func testNil() {
	var a []int

	if a == nil {
		println("i am nil")
	}
}

func makingSlices() {

}
