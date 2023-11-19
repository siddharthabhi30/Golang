package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	res := make(map[string]int)
	var arr = strings.Fields(s)
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		//extracting two values ...oen cant be undescore if it not used
		_, ok := res[arr[i]]
		// we don't really need this since if map doesn't contain the key
		// then nil val of key is chosen which is zero for int.
		if ok == false {
			res[arr[i]] = 0
		}
		res[arr[i]] = res[arr[i]] + 1

	}
	return res
}

func main() {
	fmt.Println(WordCount("hello"))
}
