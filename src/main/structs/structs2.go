package main

import "fmt"

type PointNew struct {
	x int
}

func main() {
	var structVal PointNew
	fmt.Printf("%v\n", structVal.x)
}
