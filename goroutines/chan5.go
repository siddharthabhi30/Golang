package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		//select also works for receiving ...helps in blocking calls
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci(c, quit)

	one()

	// aa:=322
	// two(aa)
	// fmt.Println(aa)
}

func one() {
	tick := make(chan int, 10)
	tick <- 10
	tick <- 20
	boom := time.After(5000 * time.Millisecond)
	for {
		fmt.Println("happen")
		select {
		case <-tick:
			fmt.Println("tick.")
			//

		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(3000 * time.Millisecond)
		}
	}
}