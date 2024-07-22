package main

/*
Delete is a No-op call if the key is not present.
*/
func main() {
	var mapp map[int]int = make(map[int]int)
	delete(mapp, 23)
}
