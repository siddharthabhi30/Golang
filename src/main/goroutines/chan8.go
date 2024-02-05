package main

import (
	"fmt"
	"sync"
	"time"
)

func one(wg *sync.WaitGroup) {
	go func() {
		fmt.Println("one  one")
		fmt.Println("one  two")
	}()

	go func() {
		fmt.Println("two  one")
		fmt.Println("two  two")
	}()
	time.Sleep(time.Second)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go one(&wg)

	wg.Wait()
}
