package main

import "fmt"

// Defining a struct type
type Address struct {
	name, city string
	Pincode    int
}

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	//two ways are possible
	"Google": {37.42202, -122.08408},
}

func main() {
	// struct literal and map literal
	add1 := Address{"Ram", "Mumbai", 221007}
	fmt.Println("Address is: ", add1)
	fmt.Println(m)

	//it needs to be done
	m := make(map[string]int)
}
