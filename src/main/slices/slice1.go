package main

import "fmt"

type Point struct {
	x int
}

/*
Slices are references not actual copy.
*/
func main() {

	testingSliceAppend()

}

func basicSliceTest() {
	var primes [3]Point
	primes[0] = Point{2}

	var primesCopy = primes[0:1]
	primesCopy[0].x = 55

	fmt.Println(primes[0].x)
	fmt.Println(primesCopy[0].x)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
How to make slice out of nowhere
*/
func makingSlice() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

/*
this proves that append makes a copy of the array.
*/
func testingSliceAppend() {
	var primes [3]Point
	firstSlice := primes[0:]
	primes2 := append(firstSlice, Point{2})
	primes2[0].x = 99
	fmt.Println(len(primes))
	fmt.Println(len(firstSlice))
	fmt.Println(len(primes2))

	fmt.Println(primes[0].x)
	fmt.Println(firstSlice[0].x)
	fmt.Println(primes2[0].x)
	//fmt.Print(primes[3].x)
}

/*
element chosen from array is a copy not reference.
if we do pointer then it will be pointer.
*/
func testingSliceElementRef() {
	var primes [3]Point
	primes[0] = Point{2}

	var chosenStruct *Point = &primes[0]
	chosenStruct.x = 99

	fmt.Println(primes[0].x)
}
