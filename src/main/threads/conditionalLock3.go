package main

import (
	"fmt"
	"sync"
	"time"
)

func listenDelayed(name string, a map[string]int, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	fmt.Println(name, " age:", a["T"])
	c.L.Unlock()
}

func broadcastFirst(name string, a map[string]int, c *sync.Cond) {
	//time.Sleep(time.Second)
	c.L.Lock()
	a["T"] = 25
	c.Broadcast()
	c.L.Unlock()
}

func main() {
	m := sync.Mutex{}
	cond := sync.NewCond(&m)
	var age = make(map[string]int)
	go listenDelayed("dsds", age, cond)
	time.Sleep(10 * time.Millisecond)
	go broadcastFirst("dsds", age, cond)
	time.Sleep(10 * time.Millisecond)
}
