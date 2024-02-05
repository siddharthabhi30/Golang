package main

import (
	crand "crypto/rand"
	"math/big"
	"os"
	"time"
)

/*
When one thread exits, it will make the whole program exit.
*/
func maybeCrash() {
	max := big.NewInt(1000)
	rr, _ := crand.Int(crand.Reader, max)
	if rr.Int64() < 330 {
		// crash!
		os.Exit(1)
	} else if rr.Int64() < 660 {
		// delay for a while.
		maxms := big.NewInt(10 * 1000)
		ms, _ := crand.Int(crand.Reader, maxms)
		time.Sleep(time.Duration(ms.Int64()) * time.Millisecond)
	}
}

func main() {

	for i := 0; i < 100; i++ {
		go maybeCrash()
	}

	time.Sleep(10 * time.Second)
}
