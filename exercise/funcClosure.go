package main

func fibonacci() func() int {
	firstEle := 0
	secondEle := 0

	cnt := 0

	return func() int {
		cnt++
		if cnt == 2 {
			return 1
		}
		if cnt == 3 {
			firstEle = 0
			secondEle = 1
		}
		secondEle = secondEle + firstEle
		firstEle = secondEle - firstEle
		return secondEle
	}
}
