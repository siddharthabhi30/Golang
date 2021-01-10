package main

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(ch chan int) {
	ch <- 23
	ch <- 23

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch <- 23
	ch <- 23
	sum1 := 0
	sum2 := 0

	for aa := range ch {
		sum1 = sum1 + aa
	}
	for aa := range ch2 {
		sum2 = sum2 + aa
	}

}

func main() {

	Same()

}
