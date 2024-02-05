package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i, j := 42, 2701

	p := &i                        // point to i
	fmt.Println(*p)                // read i through the pointer
	*p = 21                        // set i through the pointer
	fmt.Println("new val pf i", i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	var i2 *int
	i2 = &i
	*i2 = 69
	fmt.Println("this is i", *p, *i2)
	v := Vertex{1, 2}
	p2 := &v
	//no need for deferencing anything
	p2.X = 1e9
	fmt.Println(v)
	//noName := Vertex{}

	var againStruct Vertex
	fmt.Println("chrcvking this one out", againStruct)

	//this is also valid
	var noName2 Vertex
	fmt.Println("this is struct ", noName2)

}
