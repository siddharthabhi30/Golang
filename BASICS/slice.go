package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice(s)

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)

	// s = s[0:]
	// fmt.Println(s)

	check()
	slicesOfSlice()

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func check() {
	//once you increase the lower pointer you lose
	//capacity....otherwise right pointers capacity can be regained
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[2:]
	printSlice(s)
	s2 := s[0:]

	printSlice(s2)
	s3 := s2[2:]

	printSlice(s3)

}

func slicesOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)
	var pic [][]int
	var my []int
	pic = append(pic, my)
	printSlice(pic)

}
