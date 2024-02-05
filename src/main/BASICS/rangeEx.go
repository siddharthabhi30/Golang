package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	fmt.Println(res)
	for y := range res {
		res[y] = make([]uint8, dx)
		for x := range res[y] {
			res[y][x] = uint8((x + y) / 2)
		}
	}
	return res
}

func main() {
	fmt.Println(Pic(2, 3))
}
