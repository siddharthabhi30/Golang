package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fiboPusher() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
We can't close the channel and later send the message.
*/
func basicCloseTest(c chan int) {
	c <- 1
	close(c)
}

func basicCloseTestWithDelay(c chan int) {
	c <- 1
	time.Sleep(time.Second * 3)
	close(c)
}

/*
When sender closes the channel the _,ok<-chan
the ok will be false
*/
func runner() {
	c := make(chan int)
	go basicCloseTest(c)
	_, ok := <-c
	fmt.Printf("ok value gotten for the first time %v\n", ok)
	_, ok2 := <-c
	fmt.Printf("ok value gotten for the second time %v\n", ok2)
}

/*
	 When receiving from the channel the receiver
		 should have guarantee that someone will push the message to the channel
*/
func runner2() {
	c := make(chan int)
	go basicCloseTestWithDelay(c)
	for val := range c {
		fmt.Printf("Val is %v \n", val)
	}
}
func main() {
	runner2()
}
