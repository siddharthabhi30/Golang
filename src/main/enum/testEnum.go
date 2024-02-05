package main

import "fmt"

type TaskType string

const (
	Map    TaskType = "MAP"
	Reduce TaskType = "REDUCE"
	None   TaskType = "NONE"
)

func main() {
	var te TaskType
	te = None

	switch te {
	case Map:
		fmt.Print("this is map")
		break
	case Reduce:
		fmt.Println("this is reduce")
		break
	default:
		fmt.Println("unknown")
		break
	}

}
