package main

import (
	"fmt"
	"reflect"
)

func printSliceNew(s []int) {
	fmt.Printf("len=%d cap=%d %v , type %v \n", len(s), cap(s), s, reflect.TypeOf(s))
}

/*
this exp shows that the
lower index cannot go left but right index can change

although, don't use this technique.
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSliceNew(s)

	// Slice the slice to give it zero length.
	s = s[0:0]
	printSliceNew(s)

	// Extend its length.
	s = s[:4]
	printSliceNew(s)

	// Drop its first two values.
	s = s[2:4]
	printSliceNew(s)

	s = s[0:4]
	printSliceNew(s)

}
