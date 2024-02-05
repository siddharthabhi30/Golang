package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fiboSelectPusher() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
func basicSelect(c, quit chan int) {
	x := 2

	select {
	case c <- x:
		println("pushing value chan")
	case res := <-quit:
		fmt.Printf("quit recived with message %v\n", res)
		return
	}
	time.Sleep(time.Second * 5)
	// THis will never be printed because the go thread will ended by the
	// parent.
	println("this go thread ended")
}

func basicSelectPusher() {
	c := make(chan int)
	quit := make(chan int)
	go basicSelect(c, quit)

	fmt.Printf("printing receiving channel %v\n", <-c)
	//fmt.Printf("printing receiving channel %v\n", <-c)
}
func main() {
	basicSelectPusher()
}
