package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// below case doesn't work coz if main go routine ends
	// then evrything ends ..to fix above problem we can run main routine by
	//time .sleep or waiting for user input like scan....blocking call is required
	// go count("sheep")
	// go count("fish")

	var wg sync.WaitGroup
	//we are telling wait for one go routine
	wg.Add(1)
	//we are invoking function
	go func() {
		count("sheep")
		//it decrement the wg counter
		wg.Done()
	}()

	//blocking call ..will wait for wg to become done
	wg.Wait()

}

func count(thing string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
