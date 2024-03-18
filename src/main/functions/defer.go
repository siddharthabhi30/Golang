package main

import "fmt"

func task() {
	fmt.Println("deferred function")
}

func upper() int {
	defer task()
	defer println("one")
	defer println("two")
	return 33
}

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately,
//but the function call is not executed until the surrounding function returns.

// defer works like a stack in which Latest pushed function is executed.

func main() {
	res := upper()
	println(res)
}
