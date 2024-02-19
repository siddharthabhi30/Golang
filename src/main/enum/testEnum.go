package main

import "fmt"

type TaskType string

const (
	Map    TaskType = "MAppP"
	Reduce TaskType = "REDUCE"
	None   TaskType = "NONE"
)

func main() {
	var te TaskType
	te = Map

	switch te {
	case Map:
		fmt.Print("this is map")
		fmt.Printf("val is %v\n", Map)
		break
	case Reduce:
		fmt.Println("this is reduce")
		break
	default:
		fmt.Println("unknown")
		break
	}

}
