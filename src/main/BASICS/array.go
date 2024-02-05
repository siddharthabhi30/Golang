package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slicing occuring
	var s []int = primes[1:4]

	fmt.Println("len and  capacity of sliuce", len(s), cap(s))

	s[0] = 69

	//referencing like python is happening here
	fmt.Println(s, primes)

	//	this creates the same array as above,
	//	then builds a slice that references it:

	//cc:= []bool{true, true, false}

	s4 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s4)

	var s7 []int
	if s7 == nil {
		fmt.Println("is nil")
	}
}
