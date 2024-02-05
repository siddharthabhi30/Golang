package main

type Point struct {
	x int
}

func main() {
	basicMethodPassRefTest()
}

// Struct does a copy of variable not reference
func changeStruct(a Point) {
	a.x = 99

}
func basicMethodPassRefTest() {
	ss := Point{
		x: 2,
	}
	changeStruct(ss)
	print(ss.x)

}
