package main

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	ch <- 234
	if t.Left != nil {
		go Walk(t.Left, ch)
	}
	if t.Right != nil {
		go Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch)
	go Walk(t2, ch2)
	sum1 := 0
	sum2 := 0

	for {

		select {
		case msg1 := <-ch:
			sum1 = sum1 + msg1
		case msg2 := <-ch2:
			sum2 = sum2 + msg2
		}

	}

}

func main() {

	Same(tree.New(1), tree.New(1))

}
