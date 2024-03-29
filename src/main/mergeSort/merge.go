package main

import (
	"log"
	"time"
)

var arr []int

func merge(start, mid, end int) {

	start2 := mid + 1
	if arr[mid] <= arr[start2] {

		return
	}

	for start <= mid && start2 <= end {

		if arr[start] <= arr[start2] {
			start++
		} else {

			value := arr[start2]
			index := start2

			// Shift all the elements between element 1
			// element 2, right by 1.
			for index != start {
				arr[index] = arr[index-1]
				index--
			}
			arr[start] = value

			// Update all the pointers
			start++
			mid++
			start2++
		}

	}

}

func mergeSort(l, r int, ch chan int) {

	if l < r {

		// Same as (l + r) / 2, but avoids overflow
		// for large l and r
		m := l + (r-l)/2

		// Sort first and second halves
		mergeSort(l, m)
		mergeSort(m+1, r)

		merge(l, m, r)
	}
}

func test() {
	arr[0] = 69
	arr[1] = 69
}

func populate(max int) {
	for i := 1; i < max; i++ {
		arr = append(arr, max-i)
	}
}

func main() {

	start := time.Now()
	populate(100000)
	//fmt.Println(arr)
	n := len(arr)
	mergeSort(0, n-1)
	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)
	//fmt.Println(arr)

}
