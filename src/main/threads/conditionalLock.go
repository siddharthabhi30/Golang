package main

import (
	"fmt"
	"sync"
	"time"
)

/*
This shows use of broadcast
*/
func main() {

	count := 0

	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second * 2)
			mu.Lock()
			defer mu.Unlock()
			count++
			// broadcast is not a blocking statement
			cond.Broadcast()
		}()
	}

	mu.Lock()
	/*

		This for block check the condition and waits.
		cond.wait atomically releases the lock and sleeps.
	*/
	for count < 10 {
		cond.Wait()
	}

	if count == 10 {
		fmt.Println("done with all opes")
	} else {
		fmt.Printf("actual val of count %v\n", count)
	}
}
