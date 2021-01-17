package main

//https://www.dotnetperls.com/recover-go
import "fmt"

func main() {

	so := zero()
	fmt.Println(so)
}

func zero() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HERE")
			fmt.Println(err)
			return
		}
	}()
	a := 10
	b := 0
	c := a / b
	return c
}
