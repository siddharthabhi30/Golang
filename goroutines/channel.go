package main

import "fmt"

type I int

func main() {
	// go say("world")
	// say("hello")
	var c chan int
	data := 23

	//we are pushing data to c
	c <- data

	//taking data from c
	data = <-c
	fmt.Println(c)
}
