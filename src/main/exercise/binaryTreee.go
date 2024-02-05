package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, parent bool) {
	if t.Left != nil {
		Walk(t.Left, ch, false)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, false)
	}

	if parent {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	firstArr := make([]int, 0)
	secondArr := make([]int, 0)

	go Walk(t1, ch1, true)
	go Walk(t2, ch2, true)

	for val := range ch1 {
		firstArr = append(firstArr, val)
	}
	for val := range ch2 {
		secondArr = append(secondArr, val)
	}
	if len(firstArr) != len(secondArr) {
		return false
	}
	for index, _ := range firstArr {
		if firstArr[index] != secondArr[index] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
