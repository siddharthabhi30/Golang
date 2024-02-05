package main

import (
	"fmt"
	"time"
)

/*
Sender will wait for receiver to ack the message it has sent.
*/
func testChan() {
	var v chan int = make(chan int)
	//go nonRetriever()
	go retriver(v)
	println("stuck here for some seconds")
	v <- 1
	print("This line wont be printed quickly")
}

func retriver(v chan int) {
	time.Sleep(time.Second * 2)
	a, ok := <-v
	var okk int = 0
	if ok {
		okk = 1
	}
	fmt.Printf("response of val %d and ok val %d ", a, okk)
	println()

}

func nonRetriever() {
	time.Sleep(10000)
}
func main() {
	testChan()
}
