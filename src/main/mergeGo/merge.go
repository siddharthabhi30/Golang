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
		ch1 := make(chan int, 1)
		ch2 := make(chan int, 1)
		go mergeSort(l, m, ch1)
		go mergeSort(m+1, r, ch2)
		<-ch1
		<-ch2
		//fmt.Println(arr[l:m], arr[m+1:r])
		merge(l, m, r)
	}

	ch <- 1
	close(ch)

}

func populate(max int) {
	for i := 1; i < max; i++ {
		arr = append(arr, max-i)
	}
}

func main() {

	populate(600000)
	//fmt.Println(arr)
	n := len(arr)

	start := time.Now()
	ch := make(chan int, 1)
	mergeSort(0, n-1, ch)

	elapsed := time.Since(start)

	log.Printf("Binomial took %s", elapsed)

	//fmt.Println(arr)

}
