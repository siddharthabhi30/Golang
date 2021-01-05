package main

import "fmt"

func taek() {
	fmt.Println("another fuynction")
}

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

func main() {
	defer fmt.Println("world")
	taek()
	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
