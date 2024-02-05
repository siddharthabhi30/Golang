package main

import "fmt"

/*
Each caller get's its own closure variable
*/
func adderFunc() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func test_adder() {
	pos, neg := adderFunc(), adderFunc()
	print("added lines")
	println("")
	fmt.Println(pos(1))
	fmt.Println(neg(1))
}

func main() {
	test_adder()
}
