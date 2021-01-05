package main

//Go does not have classes.
//However, you can define methods on types.
import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//here dont write type
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	//we are not specifying & operator
	var pp *Vertex
	pp = &v

	//pp.scale(10) is also possible to do here
	//, Go interprets the statement v.Scale(5)
	//as (&v).Scale(5) since the Scale method has a pointer receiver.
	v.Scale(10)
	fmt.Println(v.Abs())

	Scale2(&v, 10)
	fmt.Println(v.Abs())

	var bb int = 43
	var aa *int
	aa = &bb
	fmt.Println(*aa, bb)

	var v4 Vertex
	fmt.Println(v.Abs()) // OK
	p := &v4
	fmt.Println(p.Abs()) // OK
	//p.Abs() is interpreted as (*p).Abs().
}
