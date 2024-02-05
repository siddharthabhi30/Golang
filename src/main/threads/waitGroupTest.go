package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i2 int) {
			defer wg.Done()
			fmt.Println(i2)
		}(i)
	}
	wg.Wait()

}
