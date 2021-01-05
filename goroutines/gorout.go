package main

import (
	"fmt"
	"time"
)

var n int = 23

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func do() {
	n = n + 1

}

func main() {
	// go say("world")
	// say("hello")

	for i := 0; i < 4; i++ {
		go do()
	}

	do()
	do()

	fmt.Println(n)
}
