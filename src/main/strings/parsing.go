package main

import (
	"strings"
)

func parse() {
	a := "../../mrapps/crash.so"
	fileNameWithExtension := strings.Split(a, "/")
	actualFileName := strings.Split(fileNameWithExtension[len(fileNameWithExtension)-1], ".")[0]
	return actualFileName
}

func main() {
	parse()
}
