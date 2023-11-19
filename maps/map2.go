package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	deletingFromMP()
}

func basicInitMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	testReferenceThingInMap()
}

func testReferenceThingInMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	/*
		This made a copy of the value
	*/
	tmpVertex := m["Bell Labs"]
	tmpVertex.Long = 999999

	fmt.Printf("I don't know about type %T!\n", tmpVertex)
	fmt.Println(m["Bell Labs"].Long)

	var tmpRefVertex *Vertex

	/*
			goland forbids reference to map
			because when address grows then reference to
			map will lead to dangling pointer.

		This is not the case for slice because slice
		copies the element when slice outgrows the internal
		array
	*/
	//tmpRefVertex = &(m["Bell Labs"])
	fmt.Println(tmpRefVertex.Lat)
}

func deletingFromMP() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
