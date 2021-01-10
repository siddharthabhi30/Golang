package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("sheep", c)
	// for {
	// 	//it will wait for sender to send messaage ..via some go routines...if it cant get any message from
	// 	//publisher...deadlock would happen
	// 	//to remedy this dchannel should be closed from publisher end not receiver

	// 	//TEST WHETHER CHANNEL IS OPEN OR NOT
	// 	time.Sleep(time.Millisecond * 5000)
	// 	msg, open := <-c
	// 	fmt.Println(open)
	// 	if !open {
	// 		break
	// 	}
	// 	//THIS IS printed first ..then routine thing is printed
	// 	fmt.Println("from main threaD BLOCKING ONE  ", msg)
	// }

	//it is iterating over the channel..receiving vaslue concurrently
	for msg := range c {
		//it will automatically check if the channel is closed
		fmt.Println(msg)
	}
}

//channels also have type ..they can be sent around
func count(thing string, c chan string) {
	for i := 0; i <= 5; i++ {
		//we are sending a message
		//this is blocking call for this routine
		c <- thing
		time.Sleep(time.Millisecond * 500)
		fmt.Println("from routine ", thing)

	}
	//sender should close the channel not receiver
	close(c)

}
