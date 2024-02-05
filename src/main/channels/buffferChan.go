package main

import (
	"fmt"
	"time"
)

/*
Sender will not wait for receiver to ack the message it has sent.
Since sender is in main thread, the child thread will automatically end
*/
func testBuffChan() {
	var v chan int = make(chan int, 1)
	go retriverBuff(v)
	println("stuck here for 0 seconds")
	v <- 1
	print("This line will be printed quickly")
}

/*
With buffered chan take a case
Parent-main func sent a message in already full buff
The parent will wait for receiver to pull

Receiver pulls the message and now parent is free.
Parent may close the main thread and the receiver may not
be able to do anything with the message it has received.
*/
func retriverBuff(v chan int) {
	time.Sleep(time.Second * 10)
	res := <-v
	time.Sleep(time.Second * 10)
	fmt.Println("response is %d", res)
}

func main() {
	MakingBuffChannelMOreThanFull()
}

func MakingBuffChannelMOreThanFull() {
	var v chan int = make(chan int, 1)
	go retriverBuff(v)
	println("stuck here for 0 seconds")
	v <- 1
	print("This line will be printed quickly but now stuck bcz of buffered chan is full")
	v <- 1
}
