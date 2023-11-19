package main

func main() {
	basicString()
}

/*
Read this
https://stackoverflow.com/questions/15018545/how-to-index-characters-in-a-golang-string
*/
func basicString() {
	s := "हैलो"

	s2 := "hello"

	cc := len(s)
	cc2 := len(s2)

	println(cc)
	println(cc2)

	println(s[0])
	println(s2[0])

	println(string(s[0]))
	println(string(s2[0]))

}
