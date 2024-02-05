package main

import (
	"fmt"
	"time"
)

func main() {
	one()

}

func one() {

	//we are making a buffet channel ..it wont block the sender until the buffet
	//is full
	c := make(chan string, 2)

	//deadlock coz it needs someone to receive a channel
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
	//we can push more data as expected behaviour

	c <- "hello2"
	c <- "world2"

	//we even have to close buffet channel ....deadlock
	close(c)

	//this is imp for preventing deadlock ....receiver should know that channel is closed
	for msg = range c {
		fmt.Println(msg)
	}
}
func two() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {

			c1 <- "every 500 ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {

		for {

			c2 <- "every 2 ms"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	for {
		//below code is blocking as we have to wait for c1 to finish until c2 executes
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)

		//sol fro above prob
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

}
