package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + 23
}

func add2(x int, y int) (int, int) {
	return x + 23, y + 2
}

//only to be used with sort function
func nakedReturn(x int) (x1, y int) {
	y = x + 33
	x1 = x
	return
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(3, 3))

	a, b := add2(233, 332)
	suum := add(2, 2)
	fmt.Println("suum  ", suum)

	fmt.Println(a, b)

	

}
