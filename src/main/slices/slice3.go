package main

import (
	"fmt"
	"reflect"
)

func printSliceNew(s []int) {
	fmt.Printf("len=%d cap=%d %v , type %v \n", len(s), cap(s), s, reflect.TypeOf(s))
}

func printSliceStrNew(s []string) {
	fmt.Printf("len=%d cap=%d %v , type %v \n", len(s), cap(s), s, reflect.TypeOf(s))
}

func printArrNew(s [5]int) {
	fmt.Printf("len=%d cap=%d %v , type %v \n", len(s), cap(s), s, reflect.TypeOf(s))
}

/*
Slices and array are two different thing
*/
func main() {
	copyingASliceWithNewFirstVal()
}

func basicSlice() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	printArrNew(a)
	printSliceNew(b)

	c := append(a[:], 2)
	printSliceNew(c)
}

/*
Slice when appended will double its capacity

if value to be appended are more than double, then it will

only increase upto len of extra + original len

eg: 7 will be the capacity
*/
func printAppendingOfSlice() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples", "onion", "guava"}
	food := append(veggies, fruits...)
	printSliceStrNew(food)
}

func copyingASlice() {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	printSliceStrNew(countriesCpy)
}

func copyingASliceWithNewFirstVal() {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, 0)
	countriesCpy = append(countriesCpy, "firstVal")
	countriesCpy = append(countriesCpy, neededCountries...)
	printSliceStrNew(countriesCpy)
}
