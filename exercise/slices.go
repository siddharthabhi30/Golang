package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var resSlice [][]uint8

	for i := 0; i < dy; i++ {
		resSlice = append(resSlice, make([]uint8, dx))
	}

	for i := 0; i < dy; i++ {
		for j := 0; j < dy; j++ {
			resSlice[i][j] = uint8(((i + 1) ^ (j + 1)) / 1)
		}
	}

	return resSlice
}

/*
this way is better
*/
func Pic2(dx, dy int) [][]uint8 {
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

}
