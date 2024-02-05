package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("he")

	}()
	time.Sleep(time.Millisecond * 2000)

}
