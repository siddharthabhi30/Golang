package main

import (
	"fmt"
	"time"
)

/*
any child go threads are terminated as soon as their parent gets terminated.
*/
func secondLevel() {
	fmt.Printf("this is second level \n")
}

func firstLevel() {
	//time.Sleep(20 * time.Millisecond)
	fmt.Printf("this is first level \n")
	go secondLevel()
}

func main() {
	go firstLevel()
	time.Sleep(20 * time.Millisecond)
}
