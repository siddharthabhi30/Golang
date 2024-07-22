package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)

	fmt.Printf("time elaspsed %d \n", elapsed/1000000)
}
