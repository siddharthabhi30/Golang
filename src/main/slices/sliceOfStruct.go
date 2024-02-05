package main

type Testing struct {
	a int
}

func main() {
	a := []Testing{}
	aa := make([]Testing, 1)
	if a == nil {
		println("a is nil")
	}
	aa[0].a = 99
	println(len(aa))
}
