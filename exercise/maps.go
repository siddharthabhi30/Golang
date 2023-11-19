package main

import "strings"

func WordCount(s string) map[string]int {
	var wordCounter map[string]int = make(map[string]int)
	for _, splicedArray := range strings.Fields(s) {
		wordCounter[splicedArray] += 1
	}
	return wordCounter
}

func main() {
	print(WordCount("I am learning Go!")["learning"])
}
